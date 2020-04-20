package services

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jenlesamuel/recipe-collection/httperrors"
	"github.com/jenlesamuel/recipe-collection/models"
)

//RegisterUser registers a new user
func RegisterUser(
	name string,
	email string,
	password string,
) (*models.User, *httperrors.HTTPError) {

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return user.Create()
}

//Login processes user login
func Login(email string, password string) (interface{}, *httperrors.HTTPError) {

	user, err := models.NewUserRepository().UserFromEmail(email)

	if user == nil && err == nil {
		return nil, &httperrors.HTTPError{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("User with email %s not found", email),
		}
	} else if err != nil {
		return nil, &httperrors.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: httperrors.ServerError,
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// password is invalid
		return nil, &httperrors.HTTPError{
			Status:  http.StatusUnauthorized,
			Message: "Invalid passsord",
		}
	}

	expiresAfter, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_AFTER"))

	if err != nil {
		log.Println(err)
		return nil, &httperrors.HTTPError{
			Status: http.StatusInternalServerError,
		}
	}

	tokenString, err := generateToken(
		user,
		time.Now().Unix()+int64(expiresAfter),
		os.Getenv("JWT_ISSUER"),
		os.Getenv("JWT_SECRET"),
	)

	if err != nil {
		log.Println(err.Error())

		return nil, &httperrors.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "Login Failed but its nour fault. Please try again",
		}
	}

	return tokenString, nil
}

func generateToken(user *models.User, expiresAt int64, issuer string, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": expiresAt,
		"iss": issuer,
	})

	return token.SignedString([]byte(secret))
}
