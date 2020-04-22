package modeltests

import (
	srv "ego/user/server"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = srv.Server{}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error getting env %s\n", err)
	}
	Database()

	log.Println("Before calling m.Run()!")
	retVal := m.Run()
	log.Println("After calling m.Run()!")
	os.Exit(retVal)
}

func Database() {
	var err error
	testDbName := os.Getenv("TestDbName")
	testDbDriver := os.Getenv("TestDbDriver")
	server.DB, err = gorm.Open(testDbDriver, "../"+testDbName)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", testDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", testDbDriver)
	}
	server.DB.Exec("PRAGMA foreign_keys = ON")
}

func TestVal(t *testing.T) {
	fmt.Println("testing worked!")
}
