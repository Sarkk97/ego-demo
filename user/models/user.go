package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User is a struct representing the user model
type User struct {
	ID        string     `gorm:"primary_key;varchar(120)" json:"id"`
	Phone     string     `gorm:"not null;unique" json:"phone" validate:"required,numeric"`
	Email     string     `json:"email" validate:"omitempty,email"`
	PIN       string     `gorm:"not null" json:"pin" validate:"required,numeric"`
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt time.Time  `json:"updated_at" `
	LastLogin *time.Time `json:"last_login"`
	Active    bool       `gorm:"default:true" json:"active"`
}

//UpdateUser is a struct representing the user model when updating
type UpdateUser struct {
	Phone string `json:"phone" validate:"omitempty,numeric"`
	Email string `json:"email" validate:"omitempty,email"`
}

//HashPIN hashes the entered PIN
func (u *User) HashPIN(pin string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
}

//ConfirmPIN confirms the validity of the entered pin
func (u *User) ConfirmPIN(pin string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PIN), []byte(pin))
}
