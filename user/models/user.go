package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User is a struct representing the user model
type User struct {
	ID        uint32    `gorm:"primary_key" json:"id"`
	Phone     string    `gorm:"not null;unique" json:"phone_number"`
	Email     string    `json:"email"`
	PIN       string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//HashPIN hashes the entered PIN
func (u *User) HashPIN(pin string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
}

//ConfirmPIN confirms the validity of the entered pin
func (u *User) ConfirmPIN(pin string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PIN), []byte(pin))
}
