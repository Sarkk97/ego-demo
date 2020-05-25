package repositories

import (
	"ego-api/wallet/connection"
	"ego-api/wallet/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = connection.GetDB()
	db.AutoMigrate(&models.Wallet{})
	db.AutoMigrate(&models.WalletTransaction{})
}
