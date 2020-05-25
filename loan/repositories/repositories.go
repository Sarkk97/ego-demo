package repositories

import (
	"ego-api/loan/logger"
	"ego-api/loan/models"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

//LoanRepository is a Loan collection manager
type LoanRepository interface {
	Add(loan *models.Loan) error
	GetLoanByID(id string) (*models.Loan, error)
	Update(*models.Loan) error
	Delete(*models.Loan) error
	GetUnRepaidLoan(borrowerID string) (*models.Loan, error)
}

//NewLoanRepository returns a new LoanRepository instance
func NewLoanRepository() LoanRepository {
	return &loanDBRepository{
		db:     db, //defined in init.go within same package
		logger: logger.NewLogger(),
	}
}

//LoanDBRepository is an RDBMS implementation of a LoanRepository
type loanDBRepository struct {
	db     *gorm.DB
	logger logger.Logger
}

//Add persists a loan to the db
func (repo *loanDBRepository) Add(loan *models.Loan) error {
	if err := repo.db.Create(loan).Error; err != nil {
		repo.logger.Error(err.Error())

		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}

//GetLoanById retrieves a loan by id
func (repo *loanDBRepository) GetLoanByID(id string) (*models.Loan, error) {
	loan := &models.Loan{}
	err := repo.db.Where("id = ?", id).First(loan).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		repo.logger.Error(err.Error())
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return loan, nil
}

func (repo *loanDBRepository) Update(loan *models.Loan) error {
	//check status before updating
	if err := repo.db.Model(&models.Loan{}).Updates(loan).Error; err != nil {
		repo.logger.Error(err.Error())
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}

func (repo *loanDBRepository) Delete(loan *models.Loan) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if txErr := tx.Delete(loan).Error; txErr != nil { //Delete loan
			return txErr
		}

		deletedLoan := models.NewDeletedLoan(loan)
		if txErr := tx.Create(deletedLoan).Error; txErr != nil { //Save a copy of deleted loan
			return txErr
		}

		return nil
	})

	if err != nil {
		repo.logger.Error(err.Error())

		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}

func (repo *loanDBRepository) GetUnRepaidLoan(borrowerID string) (*models.Loan, error) {
	loan := &models.Loan{}

	err := repo.db.Where("borrower_id = ? AND status <> ?", borrowerID, models.LoanStatusRepaid).First(loan).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		repo.logger.Error(err.Error())
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return loan, nil
}
