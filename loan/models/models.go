package models

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

const (
	//LoanStatusCreated is status for newly created loan
	LoanStatusCreated int = iota
	//LoanStatusPurchased is status for loan fulfilled by a lender
	LoanStatusPurchased
	//LoanStatusRepaid is status for loan repaid by the borrower
	LoanStatusRepaid
)

//Loan models a loan
type Loan struct {
	ID         string    `json:"id" gorm:"primary_key;type:varchar(120)"`
	BorrowerID string    `json:"borrowerId" gorm:"type:varchar(120);index:load_borrowerid;not null"`
	Amount     int64     `json:"amount" gorm:"type:int;not null"`
	Term       int       `json:"term" gorm:"type:int;not null"`
	Interest   float64   `json:"interest" gorm:"type:decimal;not null"`
	LenderID   string    `json:"lenderId" gorm:"type:varchar(120);index:loan_lenderid;default:null"`
	Status     int       `json:"status" gorm:"type:int;not null"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

//CalculateRepayment calulates the amount to repay for a loan
func (loan Loan) CalculateRepayment() int64 {
	//Interest amount = (interest*principal)/100 which is same as interest * (principal/100)
	interest := new(big.Rat).SetFloat64(loan.Interest)
	amount := new(big.Rat).SetInt64(loan.Amount)
	hundred := new(big.Rat).SetInt64(100)

	res := new(big.Rat)
	res = res.Mul(interest, amount).Quo(res, hundred)
	repayment, _ := res.Float64()

	return (int64(repayment) + loan.Amount)
}

//DeletedLoan is a deleted loan
type DeletedLoan struct {
	ID               string    `json:"id" gorm:"primary_key;type:varchar(120)"`
	LoanID           string    `json:"loanID" gorm:"primary_key;type:varchar(120)"`
	BorrowerID       string    `json:"borrowerId" gorm:"type:varchar(120);index:load_borrowerid;not null"`
	Amount           int64     `json:"amount" gorm:"type:int;not null"`
	Term             int       `json:"term" gorm:"type:int;not null"`
	Interest         float64   `json:"interest" gorm:"type:decimal;not null"`
	LenderID         string    `json:"lenderId" gorm:"type:varchar(120);index:loan_lenderid;default:null"`
	Status           int       `json:"status" gorm:"type:int;not null"`
	LoanCreationDate time.Time `json:"-"`
	CreatedAt        time.Time `json:"-"`
}

//NewDeletedLoan return an istance of a DeletedLoan
//derived from a Loan
func NewDeletedLoan(loan *Loan) *DeletedLoan {
	return &DeletedLoan{
		ID:               uuid.New().String(),
		LoanID:           loan.ID,
		BorrowerID:       loan.BorrowerID,
		Amount:           loan.Amount,
		Term:             loan.Term,
		Interest:         loan.Interest,
		LenderID:         loan.LenderID,
		Status:           loan.Status,
		LoanCreationDate: loan.CreatedAt,
	}
}
