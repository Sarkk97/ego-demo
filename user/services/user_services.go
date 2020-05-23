package services

import (
	"ego/user/models"
	"ego/user/repositories"
	"fmt"

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
func (s *UserService) UpdateUser(props models.UpdateUser, id string) (models.Profile, error) {
	//update user fields

	fmt.Printf("%+v\n", props)
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return models.Profile{}, err
	}
	userUpdateFields := map[string]interface{}{}
	// Update User with props
	if props.Email != "" {
		userUpdateFields["email"] = props.Email
	}
	if props.Phone != "" {
		userUpdateFields["phone"] = props.Phone
	}

	if len(userUpdateFields) != 0 {
		err := s.Repo.UpdateUser(&user, userUpdateFields)
		if err != nil {
			return models.Profile{}, err
		}
	}

	//update profile fields
	profile, err := s.Repo.GetUserProfile(user)
	if err != nil {
		return models.Profile{}, err
	}
	profileUpdateFields := map[string]interface{}{}
	// Update Profile with props
	if props.BVN != "" {
		profileUpdateFields["bvn"] = props.BVN
		//connect to BVN validation service
	}
	if props.FirstName != "" {
		profileUpdateFields["first_name"] = props.FirstName
	}
	if props.LastName != "" {
		profileUpdateFields["last_name"] = props.LastName
	}
	if props.DateOfBirth != "" {
		profileUpdateFields["date_of_birth"] = props.DateOfBirth
	}
	/*
		Handle logic for saving avatar
		if props.FirstName != "" {
			profile.FirstName = props.FirstName
		}
	*/
	if props.HomeAddress != "" {
		profileUpdateFields["home_address"] = props.HomeAddress
	}
	if props.EmploymentStatus != "" {
		profileUpdateFields["employment_status"] = props.EmploymentStatus
	}
	if props.EmployerName != "" {
		profileUpdateFields["employer_name"] = props.EmployerName
	}
	if props.EmployerAddress != "" {
		profileUpdateFields["employer_address"] = props.EmployerAddress
	}
	if props.Designation != "" {
		profileUpdateFields["designation"] = props.Designation
	}
	if props.DateOfEmployment != "" {
		profileUpdateFields["date_of_employment"] = props.DateOfEmployment
	}

	if len(profileUpdateFields) != 0 {
		err := s.Repo.UpdateUserProfile(&profile, profileUpdateFields)
		if err != nil {
			return models.Profile{}, err
		}
	}

	return profile, nil
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
