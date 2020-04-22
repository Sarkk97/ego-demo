package models

import (
	"github.com/jenlesamuel/recipe-collection/connection"
	"github.com/jinzhu/gorm"
)

//UserRepository defines methods to manage a collection of users
type UserRepository interface {
	UserFromEmail(email string) (*User, error)
	ExistWithEmail(email string) (*User, error)
}

//UserDBRepository is the RDBMS implementation of UserRepository
type UserDBRepository struct {
}

//NewUserRepository returns a user repository instance
func NewUserRepository() UserDBRepository {
	return UserDBRepository{}
}

//UserFromEmail gets a user from email
func (repo UserDBRepository) UserFromEmail(email string) (*User, error) {
	db := connection.GetDB()

	user := &User{}
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

//ExistWithEmail checks if a user with the given email exist
func (repo UserDBRepository) ExistWithEmail(email string) (bool, error) {
	db := connection.GetDB()
	user := &User{}

	if err := db.Where("email = ?", email).First(user).Error; err == nil {
		return true, nil
	} else if gorm.IsRecordNotFoundError(err) {
		return false, nil
	} else {
		return false, err
	}
}
