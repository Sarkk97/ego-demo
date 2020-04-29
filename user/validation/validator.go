package validation

import (
	"ego/user/models"

	"github.com/go-playground/validator/v10"
)

//ValidateUser is to validate user struct
func ValidateUser(u models.User) (errors map[string]string) {
	errors = map[string]string{}
	if (models.User{}) == u {
		errors["non-field"] = "The request body is empty or does not contain valid fields"
	}
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, errval := range err.(validator.ValidationErrors) {
			if errval.Tag() == "required" {
				errors[errval.Field()] = errval.Field() + " is required."
			} else {
				errors[errval.Field()] = errval.Value().(string) + " is not a valid " + errval.Tag() + " type."
			}
		}
	}
	return
}

//ValidateUpdateUser is to validate user update struct
func ValidateUpdateUser(u models.UpdateUser) (errors map[string]string) {
	errors = map[string]string{}
	if (models.UpdateUser{}) == u {
		errors["non-field"] = "The request body is empty or does not contain valid fields"
	}
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, errval := range err.(validator.ValidationErrors) {
			if errval.Tag() == "required" {
				errors[errval.Field()] = errval.Field() + " is required."
			} else {
				errors[errval.Field()] = errval.Value().(string) + " is not a valid " + errval.Tag() + " type."
			}
		}
	}
	return
}
