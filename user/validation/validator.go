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

//ValidateLoginUser is to validate user login struct
func ValidateLoginUser(u models.LoginUser) (errors map[string]string) {
	errors = map[string]string{}
	if (models.LoginUser{}) == u {
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

//ValidateRefreshTokenRequest is to validate refresh token struct
func ValidateRefreshTokenRequest(r models.RefreshToken) (errors map[string]string) {
	errors = map[string]string{}
	if (models.RefreshToken{}) == r {
		errors["non-field"] = "The request body is empty or does not contain valid fields"
	}
	validate := validator.New()
	err := validate.Struct(r)
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
