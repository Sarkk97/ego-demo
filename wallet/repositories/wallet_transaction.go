package repositories

import (
	"ego-api/wallet/httperror"
	"ego-api/wallet/models"
	"ego-api/wallet/wlogger"
	"errors"

	"github.com/jinzhu/gorm"
)

//WalletTransactionRepository is a WalletTransaction collection manager
type WalletTransactionRepository interface {
	Add(wt *models.WalletTransaction) error
	ItemWithReferenceExist(reference string) (bool, error)
}

//WalletTransactionDBRepository is an RDBMS implementation of WalletTransactionRepository
type WalletTransactionDBRepository struct {
	logger wlogger.Logger
	db     *gorm.DB
}

//NewWalletTransactionRepository instantiates a WalletTransactionRepository
func NewWalletTransactionRepository() WalletTransactionRepository {
	return &WalletTransactionDBRepository{
		logger: wlogger.NewLogger(),
		db:     db, //defined in init.go
	}
}

//Add saves a new WalletTransaction to the datastore
func (dbRepo WalletTransactionDBRepository) Add(walletTransaction *models.WalletTransaction) error {

	err := dbRepo.db.Create(walletTransaction).Error
	if err != nil {
		dbRepo.logger.Error(err.Error())

		return errors.New(httperror.ServerError)
	}

	return nil
}

//ItemWithReferenceExist checks if a WalletTransaction with
//the given reference exist
func (dbRepo *WalletTransactionDBRepository) ItemWithReferenceExist(reference string) (bool, error) {
	err := dbRepo.db.Where("paystack_reference = ?", reference).First(&models.WalletTransaction{}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		//A DB related error occurred
		dbRepo.logger.Error(err.Error())
		return false, errors.New(httperror.ServerError)
	}

	return true, nil
}
