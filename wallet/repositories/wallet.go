package repositories

import (
	"ego-api/wallet/models"
	"ego-api/wallet/wlogger"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

//WalletRepository is a Wallet collection manager
type WalletRepository interface {
	GetUserWallet(userID string) (*models.Wallet, error)
	Add(wallet *models.Wallet) error
	Update(wallet *models.Wallet) error
}

//WalletDBRepository is a database imeplementation of WalletRepository
type WalletDBRepository struct {
	db     *gorm.DB
	logger wlogger.Logger
}

//NewWalletRepository instantiates a WalletRepository
func NewWalletRepository() WalletRepository {

	return &WalletDBRepository{
		db:     db, //defined in init.go within same package
		logger: wlogger.NewLogger(),
	}
}

//GetUserWallet checks if a wallet already exist for a user
func (dbRepo *WalletDBRepository) GetUserWallet(userID string) (*models.Wallet, error) {

	wallet := &models.Wallet{}

	err := dbRepo.db.Where("user_id = ?", userID).First(wallet).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		//DB related error occurred
		dbRepo.logger.Error(err.Error())
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return wallet, nil
}

//Add saves a new Wallet instance with its associated WalletTransaction
func (dbRepo *WalletDBRepository) Add(wallet *models.Wallet) error {

	if err := dbRepo.db.Create(wallet).Error; err != nil {
		dbRepo.logger.Error(err.Error())

		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}

//Update updates the details of a wallet
func (dbRepo *WalletDBRepository) Update(wallet *models.Wallet) error {

	if err := dbRepo.db.Model(&models.Wallet{}).Updates(wallet).Error; err != nil {
		dbRepo.logger.Error(err.Error())

		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}
