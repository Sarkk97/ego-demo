package services

import (
	"ego-api/loan/data"
	"ego-api/loan/httperror"
	"ego-api/loan/logger"
	"ego-api/loan/models"
	"ego-api/loan/repositories"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

//LoanService coordinates loan related operations
type LoanService struct {
	logger         logger.Logger
	loanRepository repositories.LoanRepository
	httpClient     *HTTPClient
}

//NewLoanService Instantiates a LoanService
func NewLoanService() *LoanService {
	logger := logger.NewLogger()

	return &LoanService{
		logger:         logger,
		loanRepository: repositories.NewLoanRepository(),
		httpClient:     NewHTTPClient(),
	}
}

//RequestLoan coordinates operations for loan request
func (ls *LoanService) RequestLoan(loanData *data.LoanData) (*models.Loan, *httperror.HTTPError) {
	unpaidLoan, err := ls.loanRepository.GetUnRepaidLoan(loanData.BorrowerID)
	if err != nil { //DB error occurred
		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if unpaidLoan != nil { //Unrepaid loan exist, decline loan request
		return nil, &httperror.HTTPError{
			Message: "New loan request declined. Unrepaid loan exist",
			Code:    http.StatusBadRequest,
		}
	}

	loan := &models.Loan{
		ID:         uuid.New().String(),
		BorrowerID: loanData.BorrowerID,
		Amount:     loanData.Amount,
		Term:       loanData.Term,
		Interest:   loanData.Interest,
		LenderID:   "",
		Status:     models.LoanStatusCreated,
	}

	if err = ls.loanRepository.Add(loan); err != nil {
		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return loan, nil
}

//UpdateLoan coordinates operations for a loan update
func (ls LoanService) UpdateLoan(loanData *data.LoanData) (*models.Loan, *httperror.HTTPError) {
	loanID := loanData.LoanID
	loan, err := ls.loanRepository.GetLoanByID(loanID)
	if err != nil { //DB related error occurred
		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if loan == nil { //Invalid loan id
		return nil, &httperror.HTTPError{
			Message: "Loan does not exist",
			Code:    http.StatusNotFound,
		}
	}

	// Can only update loans that are in created state
	if loan.Status != models.LoanStatusCreated {
		return nil, &httperror.HTTPError{
			Message: "Can only update loans that have not been serviced",
			Code:    http.StatusBadRequest,
		}
	}

	loan.Amount = loanData.Amount
	loan.Term = loanData.Term
	loan.Interest = loanData.Interest

	err = ls.loanRepository.Update(loan)
	if err != nil { //DB related error occurred
		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return loan, nil
}

//DeleteLoan coordinates operations to delete a loan
func (ls *LoanService) DeleteLoan(loanID string) *httperror.HTTPError {
	loan, err := ls.loanRepository.GetLoanByID(loanID)
	if err != nil { //DB related error occurred
		return &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if loan == nil { //Invalid loan id
		return &httperror.HTTPError{
			Message: "Loan does not exist",
			Code:    http.StatusNotFound,
		}
	}

	// Can only delete loans that are in created state
	if loan.Status != models.LoanStatusCreated {
		return &httperror.HTTPError{
			Message: "Can only delete loans that has not been bought",
			Code:    http.StatusBadRequest,
		}
	}

	err = ls.loanRepository.Delete(loan)
	if err != nil {
		return &httperror.HTTPError{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//PurchaseLoan coordinates operations to purchase a loan
func (ls *LoanService) PurchaseLoan(loanID string, lenderID string) *httperror.HTTPError {
	loan, err := ls.loanRepository.GetLoanByID(loanID)
	if err != nil { //DB related error occurred
		return &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if loan == nil { //Invalid loan id
		return &httperror.HTTPError{
			Message: "Loan does not exist",
			Code:    http.StatusNotFound,
		}
	}

	if loan.Status != models.LoanStatusCreated { //only unpurchased loans can be purchased
		return &httperror.HTTPError{
			Message: "Loan already purchased",
			Code:    http.StatusBadRequest,
		}
	}

	_, httpError := ls.httpClient.transferFunds(
		lenderID,
		loan.BorrowerID,
		loan.Amount,
	)

	if httpError != nil {
		return httpError
	}

	//Update loan details
	loan.Status = models.LoanStatusPurchased
	loan.LenderID = lenderID
	err = ls.loanRepository.Update(loan)

	if err != nil { //DB related error occurred
		message := fmt.Sprintf("Purchase of loan %s by %s was successful but could not be persisted", loanID, lenderID)
		ls.logger.Error(message)

		return &httperror.HTTPError{
			Message: "An error occurred please DO NOT retry transaction. Contact site administrator",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//RepayLoan coordinates operations to repay a loan
func (ls *LoanService) RepayLoan(loanID string, borrowerID string) *httperror.HTTPError {
	loan, err := ls.loanRepository.GetLoanByID(loanID)
	if err != nil { //DB related error occurred
		return &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if loan == nil { //Invalid loan id
		return &httperror.HTTPError{
			Message: "Loan does not exist",
			Code:    http.StatusNotFound,
		}
	}

	if loan.BorrowerID != borrowerID { //Loan belongs to another borrower
		return &httperror.HTTPError{
			Message: "Cannot pay for loan belonging to another borrower",
			Code:    http.StatusBadRequest,
		}
	}

	if loan.Status != models.LoanStatusPurchased { //Loan is in created or repaid state
		return &httperror.HTTPError{
			Message: "Cannot repay unpurchased or repaid loan",
			Code:    http.StatusBadRequest,
		}
	}

	_, httpError := ls.httpClient.transferFunds(
		borrowerID,
		loan.LenderID,
		loan.CalculateRepayment(),
	)

	if httpError != nil {
		return httpError
	}

	loan.Status = models.LoanStatusRepaid
	err = ls.loanRepository.Update(loan)
	if err != nil { //DB related error occurred
		message := fmt.Sprintf("Repayment of loan %s by %s was successful but could not be persisted", loanID, borrowerID)
		ls.logger.Error(message)

		return &httperror.HTTPError{
			Message: "An error occurred please DO NOT retry transaction. Contact site administrator",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
