package factories

import (
	"ego-api/wallet/models"

	"github.com/google/uuid"
)

//WalletFactory makes new Wallet instances
type WalletFactory struct{}

//NewWalletFactory instantiates a WalletFactory
func NewWalletFactory() *WalletFactory {
	return &WalletFactory{}
}

//MakeWallet makes a new Wallet instance
func (*WalletFactory) MakeWallet(
	userID string,
	amount int) *models.Wallet {
	return &models.Wallet{
		ID:      uuid.New().String(),
		UserID:  userID,
		Balance: amount,
		//CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		//UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

//MakeWalletTransaction makes a new WalletTransaction instance
func (*WalletFactory) MakeWalletTransaction(
	walletID string,
	reference string,
	prevBalance int,
	transactionAmt int,
	currentBalance int,
	action string,
	meta string,
) *models.WalletTransaction {

	return &models.WalletTransaction{
		ID:                uuid.New().String(),
		WalletID:          walletID,
		PaystackReference: reference,
		PrevBalance:       prevBalance,
		Amount:            transactionAmt,
		Balance:           currentBalance,
		Action:            action,
		Meta:              meta,
		//CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}
