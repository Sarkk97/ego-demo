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

//NewGormRepositoryWithDB is a constructor for a gorm repo instance with a db passed into it. Useful for testing
func NewGormRepositoryWithDB(db *gorm.DB) *GormRepo {
	return &GormRepo{
		DB: db,
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

//CreateProfile creates a new profile
func (r *GormRepo) CreateProfile(p *models.Profile) error {
	err := r.DB.Create(p).Error
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

// GetUserProfile return
func (r *GormRepo) GetUserProfile(u models.User) (models.Profile, error) {
	profile := models.Profile{}
	err := r.DB.Model(&u).Related(&profile).Error
	if err != nil {
		return profile, err
	}
	profile.User = u
	return profile, nil

}

//UpdateUser updates a user
func (r *GormRepo) UpdateUser(u *models.User, fields map[string]interface{}) error {
	err := r.DB.Model(u).Updates(fields).Error
	if err != nil {
		return err
	}
	return nil
}

//UpdateUserProfile updates the user profile
func (r *GormRepo) UpdateUserProfile(p *models.Profile, fields map[string]interface{}) error {
	err := r.DB.Model(p).Updates(fields).Error
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

//UpdateUserStatus updates the user activation status
func (r *GormRepo) UpdateUserStatus(u *models.User, status bool) error {
	var err error
	err = r.DB.Model(&u).Update("active", status).Error
	if err != nil {
		return err
	}
	return nil
}
