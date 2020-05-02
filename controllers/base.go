package controllers

import (
	"ego/user/middlewares"

	"github.com/gorilla/mux"
)

//Router is the global router instance
var router = mux.NewRouter().StrictSlash(true)

//InitializeRoutes initializes the application routes
func initializeRoutes() {
	router.HandleFunc("/api/v1/user", middlewares.AuthenticationMiddleware(CreateNewUser)).Methods("POST")
	router.HandleFunc("/api/v1/users", middlewares.AuthenticationMiddleware(GetAllUsers)).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", middlewares.AuthenticationMiddleware(GetUser)).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", middlewares.AuthenticationMiddleware(UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/v1/auth/login", LoginUser).Methods("POST")
	router.HandleFunc("/api/v1/auth/refresh", RefreshToken).Methods("POST")
}

//GetRouter returns the router
func GetRouter() *mux.Router {
	initializeRoutes()
	return router
}
