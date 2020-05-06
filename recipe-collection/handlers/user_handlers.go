package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/jenlesamuel/recipe-collection/respond"
	"github.com/jenlesamuel/recipe-collection/services"
	"github.com/jenlesamuel/recipe-collection/validation"
	"github.com/thedevsaddam/govalidator"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser handlers user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	userReg := &UserReg{}

	json.NewDecoder(r.Body).Decode(userReg)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// Handle validation error
	if errorBag := userReg.validate(); !errorBag.IsEmpty() {
		respond.WithError(
			w,
			errorBag.AllErrors(),
			http.StatusBadRequest,
			headers,
		)

		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(userReg.Password),
		bcrypt.DefaultCost,
	)

	user, httpError := services.RegisterUser(
		userReg.Name,
		userReg.Email,
		string(hashedPassword),
	)

	//Handle error that occured during registration
	if httpError != nil {
		log.Println(httpError.Error()) // log error
		respond.WithError(
			w,
			httpError.Error(),
			httpError.Status,
			headers,
		)

		return
	}

	res := map[string]interface{}{
		"user": user,
	}

	//Handle success
	respond.WithSuccess(
		w,
		res,
		http.StatusCreated,
		headers,
	)
}

//UserReg is the request body for a user registration
type UserReg struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (userReg *UserReg) validate() *validation.ErrorBag {
	errorBag := &validation.ErrorBag{}

	if userReg.Name == "" {
		errorBag.Put("name", "Name is required")
	}

	if userReg.Email == "" {
		errorBag.Put("email", "Email is required")
	}

	if userReg.Password == "" {
		errorBag.Put("passowrd", "Password is required")
	}

	if userReg.ConfirmPassword == "" {
		errorBag.Put("confirmPassword", "Confirm Password is required")
	}

	return errorBag
}

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Login handles login request
func Login(w http.ResponseWriter, r *http.Request) {

	loginData := &loginBody{}
	json.NewDecoder(r.Body).Decode(loginData)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if errBag := validateLoginParams(loginData); len(errBag) > 0 {
		respond.WithError(
			w,
			errBag,
			http.StatusBadRequest,
			headers,
		)

		return
	}

	tokenString, err := services.Login(loginData.Email, loginData.Password)

	if err != nil {
		respond.WithError(
			w,
			err.Message,
			err.Status,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		map[string]string{
			"token": tokenString.(string),
		},
		http.StatusOK,
		headers,
	)

	return
}

func validateLoginParams(body *loginBody) url.Values {

	rules := govalidator.MapData{
		"email":    []string{"required", "min:6", "max:120", "email"},
		"password": []string{"required", "min:3", "max:20"},
	}

	opts := govalidator.Options{
		Data:  body,
		Rules: rules,
	}

	v := govalidator.New(opts)
	return v.ValidateStruct()
}
