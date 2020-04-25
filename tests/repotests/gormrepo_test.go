package repotests

import (
	"ego/user/models"
	"ego/user/repositories"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"gopkg.in/go-playground/assert.v1"
)

var repo repositories.GormRepo

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

	repo.DB = db
	// repo.DB.LogMode(true)

	retVal := m.Run()
	os.Exit(retVal)

}

func refreshUserTable() error {
	// db.Exec("SET foreign_key_checks=0;")
	err := repo.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	// db.Exec("SET foreign_key_checks=1;")
	err = repo.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func TestRepoCreateUser(t *testing.T) {
	var err error
	if err = refreshUserTable(); err != nil {
		t.Errorf("Can't refresh table")
	}

	user := models.User{
		Phone: "08023963212",
		Email: "testemail@gmail.com",
		PIN:   "1997",
	}

	if err = repo.CreateUser(&user); err != nil {
		t.Errorf("Error with repo: %v", err)
	}
	fmt.Printf("%v", user)
}

func TestRepoGetUsers(t *testing.T) {
	var err error
	if err = refreshUserTable(); err != nil {
		t.Errorf("Can't refresh table")
	}
	_ = repo.CreateUser(&models.User{
		Phone: "08023963212",
		Email: "testemail@gmail.com",
		PIN:   "1997",
	})

	users, err := repo.GetAllUsers()
	if err != nil {
		t.Errorf("Error with repo: %v", err)
	}

	fmt.Printf("%v", users)
	assert.Equal(t, len(users), 1)

}
