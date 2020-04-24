package models

import "github.com/jinzhu/gorm"

//Migrate is a helper function to run all db migrations
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
