package main

import (
	"ego/user/controllers"
	_ "ego/user/database"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	router := controllers.GetRouter()

	port := os.Getenv("WEB_SERVER_PORT")
	fmt.Printf("Web Server started and listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
