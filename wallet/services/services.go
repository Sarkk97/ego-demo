package services

import (
	"ego-api/wallet/data"
	"ego-api/wallet/factories"
	"ego-api/wallet/httperror"
	"ego-api/wallet/models"
	"ego-api/wallet/repositories"
	"encoding/json"
	"net/http"
)

//WalletService orchestrates the components necessary for the fulfillment
//of wallet related requests
type WalletService struct {
	httpClient            *HTTPClient
	uow                   *UnitOfWork
	factory               *factories.WalletFactory
	walletRepo            repositories.WalletRepository
	walletTransactionRepo repositories.WalletTransactionRepository
}

//NewWalletService returns a new instance of WalletService
func NewWalletService() *WalletService {

	return &WalletService{
		httpClient:            NewHTTPClient(),
		uow:                   NewUnitOfWork(),
		factory:               factories.NewWalletFactory(),
		walletRepo:            repositories.NewWalletRepository(),
		walletTransactionRepo: repositories.NewWalletTransactionRepository(),
	}

}

//FundWallet coordinates funding a wallet
func (ws *WalletService) FundWallet(
	data *data.FundWalletData) (*models.Wallet, *httperror.HTTPError) {

	//Check if transaction refernce already exist
	//An already existing transaction refernce is an attempt to
	//fund wallet with amount from an already existing transaction
	exist, err := ws.walletTransactionRepo.ItemWithReferenceExist(data.Reference)
	if err != nil {
		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if exist {
		return nil, &httperror.HTTPError{
			Message: "Transaction reference already exist",
			Code:    http.StatusBadRequest,
		}
	}

	//Verify transaction with Paystack
	res, httperr := ws.httpClient.verifyWithPaystack(data.Reference)
	if httperr != nil { // Verification failed
		return nil, httperr
	}

	status := res["status"].(bool)
	if !status {
		return nil, &httperror.HTTPError{
			Message: res["message"].(string),
			Code:    http.StatusBadRequest,
		}
	}

	verificationStatus := res["data"].(map[string]interface{})["status"].(string)
	if verificationStatus == "failed" { //Paystack payment failed
		return nil, &httperror.HTTPError{
			Message: res["data"].(map[string]interface{})["gateway_response"].(string),
			Code:    http.StatusBadRequest,
		}
	}

	amount := int(res["data"].(map[string]interface{})["amount"].(float64))

	wallet, err := ws.walletRepo.GetUserWallet(data.UserID)
	if err != nil { //DB related error occurred
		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	meta, _ := json.Marshal(res)

	if wallet == nil { //Wallet does not exist for user
		//Fund wallet by creating new wallet with associated transaction
		wallet = ws.factory.MakeWallet(data.UserID, amount)
		walletTransaction := ws.factory.MakeWalletTransaction(
			wallet.ID,
			data.Reference,
			0,
			amount,
			amount,
			models.WalletCredit,
			string(meta),
		)

		ws.uow.RegisterNew(wallet)
		ws.uow.RegisterNew(walletTransaction)

	} else {
		// Fund wallet, by updating wallet details
		wallet, err = wallet.Credit(amount)
		if err != nil {
			return nil, &httperror.HTTPError{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			}
		}

		walletTransaction := ws.factory.MakeWalletTransaction(
			wallet.ID,
			data.Reference,
			wallet.Balance-amount, //  balance before wallet was credited
			amount,
			wallet.Balance, // Current wallet balance
			models.WalletCredit,
			string(meta),
		)

		ws.uow.RegisterDirty(wallet)
		ws.uow.RegisterNew(walletTransaction)
	}

	err = ws.uow.Commit()
	if err != nil {
		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return wallet, nil
}

//TransferFunds coordinates funds transfer between wallets
func (ws *WalletService) TransferFunds(senderID string, receiverID string, amount int) *httperror.HTTPError {
	senderWallet, err := ws.walletRepo.GetUserWallet(senderID)
	if err != nil {
		return &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if senderWallet == nil { //Wallet does not exist(empty wallet)
		return &httperror.HTTPError{
			Message: "Insufficient wallet balance",
			Code:    http.StatusBadRequest,
		}
	}

	enoughFunds, err := senderWallet.HasFundsForTransaction(amount)
	if err != nil {
		return &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
	}

	if !enoughFunds { //Insufficient funds
		return &httperror.HTTPError{
			Message: "Insufficient wallet balance",
			Code:    http.StatusBadRequest,
		}
	}

	receiverWallet, err := ws.walletRepo.GetUserWallet(receiverID)
	if err != nil {
		return &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	//Debit sender wallet
	preTransferBalance := senderWallet.Balance
	senderWallet, _ = senderWallet.Debit(amount)
	senderWalletTransaction := ws.factory.MakeWalletTransaction(
		senderWallet.ID,
		"",
		preTransferBalance,
		amount,
		senderWallet.Balance,
		models.WalletDebit,
		"",
	)

	//Credit receiver wallet
	isNewReceiverWallet := false
	var preReceiptBalance int
	if receiverWallet == nil { //No wallet(empty wallet)
		isNewReceiverWallet = true
		preReceiptBalance = 0
		receiverWallet = ws.factory.MakeWallet(receiverID, amount) 
	} else {
		preReceiptBalance = receiverWallet.Balance
		receiverWallet, _ = receiverWallet.Credit(amount)
	}

	receiverWalletTransaction := ws.factory.MakeWalletTransaction(
		receiverWallet.ID,
		"",
		preReceiptBalance,
		amount,
		receiverWallet.Balance,
		models.WalletCredit,
		"",
	)

	//Persist debit and credit
	ws.uow.RegisterDirty(senderWallet)
	ws.uow.RegisterNew(senderWalletTransaction)
	if isNewReceiverWallet {
		ws.uow.RegisterNew(receiverWallet)
	} else {
		ws.uow.RegisterDirty(receiverWallet)
	}
	ws.uow.RegisterNew(receiverWalletTransaction)

	err = ws.uow.Commit()
	if err != nil {
		return &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
