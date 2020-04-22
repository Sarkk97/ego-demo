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
)

//DB is global database connection
var DB *gorm.DB

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

//Initialize initializes the DB instance
// func (conn *Conn) Initialize(dialect interface{}) {
func initialize(dialect interface{}) {
	var err error

	switch db := dialect.(type) {
	case MysqlDB:
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.DbUser, db.DbPass,
			db.DbHost, db.DbPort, db.DbName)
		DB, err = gorm.Open("mysql", DBURL)
		if err != nil {
			fmt.Println("Cannot connect to mysql database")
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to the mysql database at %s:%s\n", db.DbHost, db.DbPort)
		}
	case PostgresDB:
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", db.DbHost, db.DbPort,
			db.DbUser, db.DbName, db.DbPass)
		DB, err = gorm.Open("postgres", DBURL)
		if err != nil {
			fmt.Println("Cannot connect to postgres database")
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to the postgres database at %s:%s\n", db.DbHost, db.DbPort)
		}
	case SqliteDB:
		DB, err = gorm.Open("sqlite3", db.DbName)
		if err != nil {
			fmt.Println("Cannot connect to sqlite3 database")
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to the mysql database at %s\n", db.DbName)
		}
		DB.Exec("PRAGMA foreign_keys = ON")
	default:
		log.Fatal("Unrecognized Database type")
	}

	//perform models migration here
	models.Migrate(DB)

}

//Run is to run the DB setup
func Run() {

	dbparams := MysqlDB{
		DbHost: os.Getenv("DB_HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbPass: os.Getenv("DB_PASSWORD"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
	}

	//for sqlite
	// db := SqliteDB{
	// 	DbName: os.Getenv("DB_NAME")
	// }

	initialize(dbparams)
}
