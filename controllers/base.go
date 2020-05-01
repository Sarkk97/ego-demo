package controllers

import (
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
	router.HandleFunc("/api/v1/login", LoginUser).Methods("POST")
}

//GetRouter returns the router
func GetRouter() *mux.Router {
	initializeRoutes()
	return router
}
