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
)

//LoginUser handler logs in user
func LoginUser(w http.ResponseWriter, r *http.Request) {
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
	user := models.LoginUser{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}
	//validate request body
	errs := valid.ValidateLoginUser(user)
	if len(errs) != 0 {
		response.Error(w, errs, 400, headers)
		return
	}
	//get UserService
	service := services.NewUserService(repositories.NewGormRepository())

	//service to login user
	tokens, err := service.LoginUser(user)
	if err != nil {
		response.Error(w, err.Error(), 400, headers)
		return
	}
	response.Success(w, tokens, 200, headers)
}
