package controllers

import (
	"ego/user/models"
	"ego/user/repositories"
	"ego/user/response"
	"ego/user/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
		response.Error(w, err, 400, headers)
		return
	}

	//marshal request body into json
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, err, 400, headers)
		return
	}

	//validate the request body
	errs := user.Validate()
	if len(errs) != 0 {
		response.Error(w, errs, 400, headers)
		return
	}
	//get UserService
	service := services.NewUserService(repositories.NewGormRepository())
	//service to create user
	err = service.CreateUser(&user)
	if err != nil {
		response.Error(w, err, 400, headers)
		return
	}
	response.Success(w, user, 201, headers)

}
