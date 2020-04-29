package repositories

import (
	"ego/user/database"
	"ego/user/models"

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
	err := r.DB.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

//GetAllUsers gets all users
func (r *GormRepo) GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	err := r.DB.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

//GetUser gets a specific user
func (r *GormRepo) GetUser(id string) (models.User, error) {
	user := models.User{}
	err := r.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

//UpdateUser updates a user
func (r *GormRepo) UpdateUser(u models.UpdateUser, id string) (models.User, error) {
	user := models.User{}
	err := r.DB.Model(&user).Where("id = ?", id).Updates(u).Error
	if err != nil {
		return models.User{}, err
	}
	updatedUser, _ := r.GetUser(id)
	return updatedUser, nil
}
