package models

import (
	"log"
	"net/http"

	"github.com/jenlesamuel/recipe-collection/httperrors"

	"github.com/jenlesamuel/recipe-collection/connection"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = connection.GetDB()
	db.AutoMigrate(&User{})
}

//User represents a user type
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(120);not null;unique"`
	Password string `json:"-" gorm:"type:varchar(120); not null"`
}

//Create creates a user
func (user *User) Create() (*User, *httperrors.HTTPError) {
	exist, err := NewUserRepository().ExistWithEmail(user.Email)

	if err != nil {
		// Error occurred while trying to check if user with email already exist
		return nil, &httperrors.HTTPError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if exist == true {
		return nil, &httperrors.HTTPError{
			Status:  http.StatusBadRequest,
			Message: "User with email already exist",
		}
	}

	db := connection.GetDB()

	if err := db.Create(user).Error; err != nil {
		log.Println(err.Error())
		return nil, &httperrors.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: httperrors.ServerError,
		}
	}

	return user, nil
}
