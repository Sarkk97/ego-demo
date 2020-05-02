package services

import (
	"ego/user/models"
	"ego/user/repositories"

	"github.com/google/uuid"
)

//UserService is a struct for user services
type UserService struct {
	// Repo *repositories.GormRepo
	Repo repositories.UserRepository
}

//NewUserService is constructor for UserService
func NewUserService(r repositories.UserRepository) *UserService {
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
func (s *UserService) UpdateUser(props models.UpdateUser, id string) (models.User, error) {
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	// Update User with props
	if props.Email != "" {
		user.Email = props.Email
	}
	if props.Phone != "" {
		user.Phone = props.Phone
	}

	err = s.Repo.UpdateUser(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
