package tests

import (
	"ego/user/models"
	"ego/user/repositories"
	"ego/user/services"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var repo repositories.UserRepository
var service *services.UserService

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %s\n", err)
	}
	Database()
	repo = repositories.NewGormRepositoryWithDB(db)
	service = services.NewUserService(repo)
	// db.LogMode(true)
	retVal := m.Run()
	os.Exit(retVal)
}

func Database() {
	var err error
	testDbName := os.Getenv("TestDbName")
	testDbDriver := os.Getenv("TestDbDriver")
	db, err = gorm.Open(testDbDriver, testDbName)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", testDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", testDbDriver)
	}
	db.Exec("PRAGMA foreign_keys = ON")
}

func refreshUserTable() error {
	// db.Exec("SET foreign_key_checks=0;")
	err := db.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	// db.Exec("SET foreign_key_checks=1;")
	err = db.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func seedOneUser() (models.User, error) {
	user := models.User{
		ID:    uuid.New().String(),
		Phone: "08023963212",
		Email: "testemail@gmail.com",
		PIN:   "1997",
	}
	err := db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() error {
	users := []models.User{
		{
			ID:    uuid.New().String(),
			Phone: "08023963212",
			Email: "testemail@gmail.com",
			PIN:   "1997",
		},
		{
			ID:    uuid.New().String(),
			Phone: "08023863412",
			Email: "testemail2@gmail.com",
			PIN:   "1999",
		},
	}

	for i := range users {
		err := db.Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}
