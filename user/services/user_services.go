package services

import (
	"ego/user/models"
	"ego/user/repositories"
)

//UserService is a struct for user services
type UserService struct {
	Repo *repositories.GormRepo
}

//NewUserService is constructor for UserService
func NewUserService(r *repositories.GormRepo) *UserService {
	return &UserService{
		Repo: r,
	}
}

//CreateUser is a user service method to create a user
func (s *UserService) CreateUser(u *models.User) (err error) {
	//Hash User PIN
	hashedPIN, err := u.HashPIN(u.PIN)
	if err != nil {
		return err
	}

	u.PIN = string(hashedPIN)
	err = s.Repo.CreateUser(u)
	if err != nil {
		return err
	}
	//Send Email notification after creation
	// go send_email(u.Email)
	return nil
}
