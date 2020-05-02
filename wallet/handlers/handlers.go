package handlers

import (
	"ego-api/wallet/data"
	"ego-api/wallet/respond"
	"ego-api/wallet/services"
	"encoding/json"
	"net/http"
)

//Index is the handler for the base route
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Welcome: EGO wallet api v1")
}

//FundWallet credits a wallet
func FundWallet(w http.ResponseWriter, r *http.Request) {
	requestData := &data.FundWalletData{}
	json.NewDecoder(r.Body).Decode(requestData)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if errBag := requestData.Validate(); len(errBag) != 0 { // Validation error occurred
		respond.WithError(
			w,
			errBag,
			400,
			headers,
		)

		return
	}

	walletService := services.NewWalletService()
	wallet, httperr := walletService.FundWallet(requestData)
	if httperr != nil {
		respond.WithError(
			w,
			httperr.Message,
			httperr.Code,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		*wallet,
		http.StatusCreated,
		headers,
	)

	return
}

//TransferFunds handles funds transfer
func TransferFunds(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	senderID := vars["senderID"]
	receiverID := vars["receiverID"]

	requestBody := &struct {
		amount int `json:"amount"`
	}

	json.NewDecoder(r.Body).Decode(requestBody)*/

	//TODO validate amount

}
