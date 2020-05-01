package services

import (
	"ego/user/auth"
	"ego/user/models"
	"ego/user/repositories"

	"github.com/google/uuid"
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
	u.ID = uuid.New().String()
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

//GetUsers is a user service method to get all user
func (s *UserService) GetUsers() (users []models.User, err error) {
	return s.Repo.GetAllUsers()
}

//GetUser is a user service method to get a user
func (s *UserService) GetUser(id string) (models.User, error) {
	return s.Repo.GetUser(id)
}

//UpdateUser is a user service method to update a user
func (s *UserService) UpdateUser(user models.UpdateUser, id string) (models.User, error) {
	return s.Repo.UpdateUser(user, id)
}

//LoginUser is a user service method to login a user
func (s *UserService) LoginUser(user models.LoginUser) (map[string]string, error) {
	loggedInUser, err := s.Repo.LoginUser(user)
	tokens := map[string]string{}
	if err != nil {
		return tokens, err
	}
	//create user tokens
	tokens, err = auth.CreateTokens(loggedInUser.ID)
	if err != nil {
		return tokens, err
	}
	//update user last_login
	_ = s.Repo.UpdateUserLogin(loggedInUser)
	return tokens, nil
}
