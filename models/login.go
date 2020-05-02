package models

//LoginUser is a struct representing the user model during login
type LoginUser struct {
	Phone string `json:"phone" validate:"required,numeric"`
	PIN   string `json:"pin" validate:"required,numeric"`
}

//RefreshToken is a struct representing the refresh token
type RefreshToken struct {
	Refresh string `json:"refresh" validate:"required"`
}
