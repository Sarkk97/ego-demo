package main

import (
	"ego/user/controllers"
	_ "ego/user/database"
)

func main() {

	//load the environment variable

	//db init and migrations
	// database.Run()

	//webserver init and routes initialization
	controllers.Run()
}
