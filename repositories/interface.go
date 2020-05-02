package repositories

import "ego/user/models"

//UserRepository is an interface to be implemented by different storage layers
type UserRepository interface {
	CreateUser(*models.User) error
	GetAllUsers() ([]models.User, error)
	GetUser(string) (models.User, error)
	UpdateUser(*models.User) error
	LoginUser(models.LoginUser) (models.User, error)
}
