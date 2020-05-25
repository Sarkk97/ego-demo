package repositories

import (
	"ego-api/loan/connection"
	"ego-api/loan/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = connection.GetDB()
	db.AutoMigrate(&models.Loan{})
	db.AutoMigrate(&models.DeletedLoan{})
}
