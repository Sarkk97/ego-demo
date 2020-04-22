package main

import (
	"ego/user/controllers"
	"ego/user/database"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	//load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	} else {
		fmt.Println("Environment set successfully")
	}

	//db init and migrations
	database.Run()

	//webserver init and routes initialization
	controllers.Run()
}
