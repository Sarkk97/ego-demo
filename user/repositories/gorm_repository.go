package repositories

import (
	"ego/user/models"

	"github.com/jinzhu/gorm"
)

//GormRepo is a repository for gorm orm
type GormRepo struct {
	DB *gorm.DB
}

//NewGormRepository is a constructor for a gorm repo instance
func NewGormRepository(db *gorm.DB) *GormRepo {
	return &GormRepo{
		DB: db,
	}
}

//CreateUser creates a new user
func (r *GormRepo) CreateUser(u *models.User) (err error) {
	err = r.DB.Create(&u).Error
	return
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
