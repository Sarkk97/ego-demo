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
	//create profile
	profile := models.Profile{
		ID:   uuid.New().String(),
		User: *u,
	}
	err = s.Repo.CreateProfile(&profile)
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

//GetUserProfile is a user service method to get a user profile
func (s *UserService) GetUserProfile(id string) (models.Profile, error) {
	// profile, err := s.Repo.GetProfile(id)
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return models.Profile{}, err
	}
	// fmt.Printf("%+v\n", profile)
	profile, err := s.Repo.GetUserProfile(user)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
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

//UserActivation is a user service method to activate/deactivate a user
func (s *UserService) UserActivation(id string, status bool) (models.User, error) {
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	err = s.Repo.UpdateUserStatus(&user, status)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}
