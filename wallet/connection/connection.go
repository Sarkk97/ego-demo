package connection

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Load mysql driver
)

var db *gorm.DB

func init() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)

	conn, err := gorm.Open("mysql", connString)

	if err != nil {
		log.Fatalln("Could not establish connection to database")
	}

	db = conn
}

//GetDB returns a GORM database connection instance
func GetDB() *gorm.DB {
	return db
}
