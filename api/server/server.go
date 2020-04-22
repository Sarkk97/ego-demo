package server

//This file is to actually run the server

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var server = Server{}

//Run is to start the server
func Run() {

	//load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	} else {
		fmt.Println("We are getting env values")
	}

	db := MysqlDB{
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

	server.Initialize(db)
	server.Run(":8080")

}
