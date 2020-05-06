package controllers

import (
	"ego/user/middlewares"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Router is the global router instance
var router = mux.NewRouter()

//InitializeRoutes initializes the application routes
func initializeRoutes() {
	router := router.PathPrefix("/api/v1").Subrouter()
	AuthHandlers := alice.New(middlewares.TerminalLoggingHandler, middlewares.AuthenticationMiddleware)
	NonAuthHandlers := alice.New(middlewares.TerminalLoggingHandler)

	router.Handle("/user", AuthHandlers.ThenFunc(CreateNewUser)).Methods("POST")
	router.Handle("/users", AuthHandlers.ThenFunc(GetAllUsers)).Methods("GET")
	router.Handle("/user/{id}", AuthHandlers.ThenFunc(GetUser)).Methods("GET")
	router.Handle("/user/{id}", AuthHandlers.ThenFunc(UpdateUser)).Methods("PUT")
	router.Handle("/auth/login", NonAuthHandlers.ThenFunc(LoginUser)).Methods("POST")
	router.Handle("/auth/refresh", NonAuthHandlers.ThenFunc(RefreshToken)).Methods("POST")
}

//GetRouter returns the router
func GetRouter() *mux.Router {
	initializeRoutes()
	return router
}
