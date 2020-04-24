package models

import (
	"time"
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
