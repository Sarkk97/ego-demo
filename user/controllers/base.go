package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Router is the global router instance
var router = mux.NewRouter().StrictSlash(true)

//InitializeRoutes initializes the application routes
func initializeRoutes() {

}

//GetRouter returns the router
func GetRouter() *mux.Router {
	return router
}

//Run is to run the router setup
func Run() {
	initializeRoutes()
	port := os.Getenv("WEB_SERVER_PORT")
	fmt.Printf("Web Server started and listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
