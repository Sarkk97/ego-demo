package controllers

import (
	"ego/user/models"
	"ego/user/repositories"
	"ego/user/response"
	"ego/user/services"
	valid "ego/user/validation"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

/**
For a standard,

Post handlers request body must be validated,
There should be a standard response format,
**/

//CreateNewUser handler creates a new user
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}

	//marshal request body into json
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}

	//validate the request body
	errs := valid.ValidateUser(user)
	if len(errs) != 0 {
		response.Error(w, errs, 400, headers)
		return
	}
	//get UserService
	service := services.NewUserService(repositories.NewGormRepository())
	//service to create user
	err = service.CreateUser(&user)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}
	response.Success(w, user, 201, headers)

}

//GetAllUsers handler gets all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	//get UserService
	service := services.NewUserService(repositories.NewGormRepository())
	//service to get all users
	users, err := service.GetUsers()
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}
	response.Success(w, users, 200, headers)
}

//GetUser handler gets all users
func GetUser(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	vars := mux.Vars(r)
	id := vars["id"]

	//get UserService
	service := services.NewUserService(repositories.NewGormRepository())
	//service to get user
	user, err := service.GetUser(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.Error(w, err.Error(), 404, headers)
		} else {
			response.Error(w, err.Error(), 400, headers)
		}
		return
	}
	response.Success(w, user, 200, headers)
}

//GetUserProfile handler gets user profile
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	vars := mux.Vars(r)
	id := vars["id"]

	//get UserService
	service := services.NewUserService(repositories.NewGormRepository())
	//service to get user
	profile, err := service.GetUserProfile(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.Error(w, err.Error(), 404, headers)
		} else {
			response.Error(w, err.Error(), 400, headers)
		}
		return
	}
	response.Success(w, profile, 200, headers)
}

//UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	service := services.NewUserService(repositories.NewGormRepository())

	vars := mux.Vars(r)
	id := vars["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}

	//marshal request body into json
	userProps := models.UpdateUser{}
	err = json.Unmarshal(body, &userProps)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}

	//validate request body
	errs := valid.ValidateUpdateUser(userProps)
	if len(errs) != 0 {
		response.Error(w, errs, 400, headers)
		return
	}

	updatedUser, err := service.UpdateUser(userProps, id)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}
	response.Success(w, updatedUser, 200, headers)

}

//UserActivation toggles user activation status
func UserActivation(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	service := services.NewUserService(repositories.NewGormRepository())
	vars := mux.Vars(r)
	id := vars["id"]
	action := vars["action"]
	user := models.User{}
	var err error

	if action == "activate" {
		user, err = service.UserActivation(id, true)
	} else {
		user, err = service.UserActivation(id, false)
	}

	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}
	response.Success(w, user, 200, headers)

}
