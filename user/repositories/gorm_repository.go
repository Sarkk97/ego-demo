package repositories

import (
	"ego/user/database"
	"ego/user/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

//GormRepo is a repository for gorm orm
type GormRepo struct {
	DB *gorm.DB
}

//NewGormRepository is a constructor for a gorm repo instance
func NewGormRepository() *GormRepo {
	return &GormRepo{
		DB: database.GetDB(),
	}
}

//CreateUser creates a new user
func (r *GormRepo) CreateUser(u *models.User) error {
	err := r.DB.Debug().Create(u).Error
	if err != nil {
		return err
	}
	fmt.Printf("repo: %v\n", u)
	return nil
}

//GetAllUsers gets all users
func (r *GormRepo) GetAllUsers() (users []models.User, err error) {
	err = r.DB.Find(&users).Error
	if err != nil {
		users = []models.User{}
		return
	}
	return
}
