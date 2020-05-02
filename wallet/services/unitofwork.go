package services

import (
	"ego-api/wallet/connection"
	"ego-api/wallet/models"
	"ego-api/wallet/repositories"

	"github.com/jinzhu/gorm"
)

//UnitOfWork is an implementation of the unit of work design pattern
type UnitOfWork struct {
	new                   []interface{}
	dirty                 []interface{}
	walletRepo            repositories.WalletRepository
	walletTransactionRepo repositories.WalletTransactionRepository
	db                    *gorm.DB
}

//NewUnitOfWork instantiates a UnitOfWork
func NewUnitOfWork() *UnitOfWork {
	return &UnitOfWork{
		new:                   []interface{}{},
		dirty:                 []interface{}{},
		walletRepo:            repositories.NewWalletRepository(),
		walletTransactionRepo: repositories.NewWalletTransactionRepository(),
		db:                    connection.GetDB(),
	}
}

//RegisterNew registers a new object to be Created
func (uow *UnitOfWork) RegisterNew(obj interface{}) {
	uow.new = append(uow.new, obj)
}

//RegisterDirty registers a new object to be updated
func (uow *UnitOfWork) RegisterDirty(obj interface{}) {
	uow.dirty = append(uow.dirty, obj)
}

//Commit carries out the actual persistence of objects to the datastore
func (uow *UnitOfWork) Commit() error {

	err := uow.db.Transaction(func(tx *gorm.DB) error {
		for _, new := range uow.new { // Save new objects

			switch obj := new.(type) {
			case *models.Wallet:
				err := uow.walletRepo.Add(obj)
				if err != nil {
					return err
				}

			case *models.WalletTransaction:
				err := uow.walletTransactionRepo.Add(obj)
				if err != nil {
					return err
				}
			}
		}

		for _, dirty := range uow.dirty { //Update objects
			switch obj := dirty.(type) {
			case *models.Wallet:
				err := uow.walletRepo.Update(obj)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		// err is already logged in repositories
		return err
	}

	return nil
}
