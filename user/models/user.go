package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

//User is a struct representing the user model
type User struct {
	ID        uint32    `gorm:"primary_key" json:"id"`
	Phone     string    `gorm:"not null;unique" json:"phone" validate:"required,numeric"`
	Email     string    `json:"email" validate:"omitempty,email"`
	PIN       string    `gorm:"not null" json:"pin" validate:"required,numeric"`
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
}

//HashPIN hashes the entered PIN
func (u *User) HashPIN(pin string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
}

//ConfirmPIN confirms the validity of the entered pin
func (u *User) ConfirmPIN(pin string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PIN), []byte(pin))
}

//Validate is to validate user struct
func (u *User) Validate() (errors map[string]string) {
	errors = map[string]string{}
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, errval := range err.(validator.ValidationErrors) {
			if errval.Tag() == "required" {
				errors[errval.Field()] = errval.Field() + " is required."
			} else {
				errors[errval.Field()] = errval.Value().(string) + " is not a valid " + errval.Tag() + " type."
			}
		}
	}
	return
}
