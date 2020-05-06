package connection

import (
	"fmt"
	"log"
	"os"

	"github.com/jenlesamuel/recipe-collection/helpers"

	"github.com/jinzhu/gorm"
	//Load mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {

	helpers.LoadENV()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)

	conn, err := gorm.Open("mysql", connString)

	if err != nil {
		log.Fatal("Could not establish connection to database")
	}

	db = conn
}

//GetDB returns a database connection instance
func GetDB() *gorm.DB {
	return db
}
