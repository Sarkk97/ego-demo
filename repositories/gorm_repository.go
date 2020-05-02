package repositories

import (
	"ego/user/database"
	"ego/user/models"
	"time"

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
func (r *GormRepo) UpdateUser(u *models.User) error {
	err := r.DB.Save(u).Error
	if err != nil {
		return err
	}
	return nil
}

//LoginUser validates a login and returns the user
func (r *GormRepo) LoginUser(u models.LoginUser) (models.User, error) {
	var err error
	//check if user exists
	user := models.User{}
	err = r.DB.First(&user, "phone = ?", u.Phone).Error
	if err != nil {
		return models.User{}, err
	}
	//confirm the pin
	err = user.ConfirmPIN(u.PIN)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

//UpdateUserLogin updates the user last_login timestamp after a successful login
func (r *GormRepo) UpdateUserLogin(u models.User) error {
	var err error
	err = r.DB.Model(&u).UpdateColumn("last_login", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
