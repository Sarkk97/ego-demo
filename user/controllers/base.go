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
	router.HandleFunc("/api/v1/user", CreateNewUser).Methods("POST")
	router.HandleFunc("/api/v1/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", UpdateUser).Methods("PUT")
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
