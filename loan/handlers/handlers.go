package handlers

import (
	"ego-api/loan/data"
	"ego-api/loan/respond"
	"ego-api/loan/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	specError = "Request parameters do not match specification"
)

//RequestLoan handles loan request
func RequestLoan(w http.ResponseWriter, r *http.Request) {
	loanData := &data.LoanData{}
	err := json.NewDecoder(r.Body).Decode(loanData)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if err != nil { //Request paramers do not match specification
		respond.WithError(
			w,
			specError,
			http.StatusUnprocessableEntity,
			headers,
		)

		return
	}

	if errBag := loanData.Validate(data.PostDataType); len(errBag) != 0 {
		respond.WithError(
			w,
			errBag,
			http.StatusBadRequest,
			headers,
		)

		return
	}

	loan, httpError := services.NewLoanService().RequestLoan(loanData)
	if httpError != nil {
		respond.WithError(
			w,
			httpError.Message,
			httpError.Code,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		loan,
		http.StatusCreated,
		headers,
	)
}

//UpdateLoan handles request to update a loan
func UpdateLoan(w http.ResponseWriter, r *http.Request) {
	loanID := mux.Vars(r)["LoanID"]
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	loanData := &data.LoanData{}
	err := json.NewDecoder(r.Body).Decode(loanData)
	if err != nil {
		respond.WithError(
			w,
			specError,
			http.StatusUnprocessableEntity,
			headers,
		)
		return
	}

	loanData.LoanID = loanID
	if errBag := loanData.Validate(data.PutDataType); len(errBag) != 0 {
		respond.WithError(
			w,
			errBag,
			http.StatusBadRequest,
			headers,
		)
		return
	}

	loan, httpError := services.NewLoanService().UpdateLoan(loanData)
	if httpError != nil {
		respond.WithError(
			w,
			httpError.Error,
			httpError.Code,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		loan,
		http.StatusOK,
		headers,
	)
}

//DeleteLoan handles request to delete a loan
func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	//TODO: Write middle ware to validate that client issuing request is resource owner
	loanID := mux.Vars(r)["LoanID"]
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	httpError := services.NewLoanService().DeleteLoan(loanID)
	if httpError != nil {
		respond.WithError(
			w,
			httpError.Message,
			httpError.Code,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		"",
		http.StatusNoContent,
		headers,
	)
}

//PurchaseLoan handles request tonpurchasea loan
func PurchaseLoan(w http.ResponseWriter, r *http.Request) {
	routeParams := mux.Vars(r)
	loanID := routeParams["LoanID"]
	lenderID := routeParams["LenderID"]
	headers := map[string]string{
		"Content_Type": "application/json",
	}

	httpError := services.NewLoanService().PurchaseLoan(loanID, lenderID)
	if httpError != nil {
		respond.WithError(
			w,
			httpError.Message,
			httpError.Code,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		http.StatusText(http.StatusOK),
		http.StatusOK,
		headers,
	)
}

//RepayLoan handles request to repay a loan
func RepayLoan(w http.ResponseWriter, r *http.Request) {
	routeParameters := mux.Vars(r)
	loanID := routeParameters["LoanID"]
	borrowerID := routeParameters["BorrowerID"]
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	httpError := services.NewLoanService().RepayLoan(loanID, borrowerID)
	if httpError != nil {
		respond.WithError(
			w,
			httpError.Message,
			httpError.Code,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		http.StatusText(http.StatusOK),
		http.StatusOK,
		headers,
	)
}
