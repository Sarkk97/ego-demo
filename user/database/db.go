package database

//This file contains all function invocations to run the server

import (
	"fmt"
	"log"
	"os"

	"ego/user/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgress database driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //sqlite database driver
	"github.com/joho/godotenv"
)

//DB is global database connection
var dbconn *gorm.DB

//MysqlDB struct for mysql connection params
type MysqlDB struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

//PostgresDB struct for postgres connection params
type PostgresDB struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

//SqliteDB struct for sqlite connection params
type SqliteDB struct {
	DbName string
}

var dialect = MysqlDB{
	DbHost: os.Getenv("DB_HOST"),
	DbName: os.Getenv("DB_NAME"),
	DbPass: os.Getenv("DB_PASSWORD"),
	DbPort: os.Getenv("DB_PORT"),
	DbUser: os.Getenv("DB_USER"),
}

//Init initializes the DB instance
// func (conn *Conn) Initialize(dialect interface{}) {
func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	} else {
		fmt.Println("Environment set successfully")
	}

	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")
	DbPass := os.Getenv("DB_PASSWORD")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPass,
		DbHost, DbPort, DbName)
	dbconn, err = gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Println("Cannot connect to mysql database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("Connected to the mysql database at %s:%s\n", DbHost, DbPort)
	}

	//perform models migration here
	dbconn.LogMode(true)
	models.Migrate(dbconn)

}

//GetDB returns the database connection
func GetDB() *gorm.DB {
	return dbconn
}
