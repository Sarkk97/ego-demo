package servicestests

import (
	"ego/user/models"
	"ego/user/repositories"
	"ego/user/services"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"gopkg.in/go-playground/assert.v1"
)

var service *services.UserService

func TestMain(m *testing.M) {
	//initialize testdb
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error getting env %s\n", err)
	}
	db, err := gorm.Open(os.Getenv("TestDbDriver"), "../"+os.Getenv("TestDbName"))
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", os.Getenv("TestDbDriver"))
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", os.Getenv("TestDbDriver"))
	}
	db.Exec("PRAGMA foreign_keys = ON")

	repo := repositories.NewGormRepository(db)
	service = services.NewUserService(repo)
	// repo.DB.LogMode(true)

	retVal := m.Run()
	os.Exit(retVal)

}

func refreshUserTable() error {
	// db.Exec("SET foreign_key_checks=0;")
	err := service.Repo.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	// db.Exec("SET foreign_key_checks=1;")
	err = service.Repo.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func TestUserCreateService(t *testing.T) {
	var err error
	if err = refreshUserTable(); err != nil {
		t.Errorf("Error refreshing table: %v", err)
	}

	user := models.User{
		Phone: "08023963212",
		Email: "testemail@gmail.com",
		PIN:   "1997",
	}

	err = service.CreateUser(&user)
	if err != nil {
		t.Errorf("Error with service: %v", err)
	}

	assert.Equal(t, user.ID, uint32(1))
	if err = user.ConfirmPIN("1997"); err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}
}
