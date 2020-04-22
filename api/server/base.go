package server

//This file contains all function invocations to run the server

import (
	"fmt"
	"log"
	"net/http"

	"ego/api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgress database driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //sqlite database driver
)

//Server struct to hold global server connection
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

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

//Server methods

//Initialize initializes the server instance with DB and Router credentials
func (s *Server) Initialize(dialect interface{}) {
	var err error

	switch db := dialect.(type) {
	case MysqlDB:
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.DbUser, db.DbPass,
			db.DbHost, db.DbPort, db.DbName)
		s.DB, err = gorm.Open("mysql", DBURL)
		if err != nil {
			fmt.Println("Cannot connect to mysql database")
			log.Fatal("This is the error:", err)
		} else {
			fmt.Println("We are connected to the mysql database")
		}
	case PostgresDB:
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", db.DbHost, db.DbPort,
			db.DbUser, db.DbName, db.DbPass)
		s.DB, err = gorm.Open("postgres", DBURL)
		if err != nil {
			fmt.Println("Cannot connect to postgres database")
			log.Fatal("This is the error:", err)
		} else {
			fmt.Println("We are connected to the postgres database")
		}
	case SqliteDB:
		s.DB, err = gorm.Open("sqlite3", db.DbName)
		if err != nil {
			fmt.Println("Cannot connect to sqlite3 database")
			log.Fatal("This is the error:", err)
		} else {
			fmt.Println("We are connected to the sqlite3 database")
		}
		s.DB.Exec("PRAGMA foreign_keys = ON")
	default:
		log.Fatal("Unrecognized Database type")
	}

	//perform models migration here
	models.Migrate(s.DB)

	//initialize routers here

}

//Run runs the server on the passed in port address
func (s *Server) Run(addr string) {
	fmt.Printf("Listening on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
