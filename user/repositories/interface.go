package repositories

import "ego/user/models"

//UserRepository is an interface to be implemented by different storage layers
type UserRepository interface {
	CreateUser(*models.User) error
	CreateProfile(*models.Profile) error
	GetAllUsers() ([]models.User, error)
	GetUser(string) (models.User, error)
	UpdateUser(*models.User, map[string]interface{}) error
	LoginUser(models.LoginUser) (models.User, error)
	UpdateUserStatus(*models.User, bool) error
	// GetProfile(string) (models.Profile, error)
	GetUserProfile(models.User) (models.Profile, error)
	UpdateUserProfile(*models.Profile, map[string]interface{}) error
}
