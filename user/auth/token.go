package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateAccessToken generates the JWT access token with claims
func CreateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["token_type"] = "access"
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

//CreateRefreshToken generates the JWT refresh token with claims
func CreateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["token_type"] = "refresh"
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix() //Token expires after 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

//CreateTokens generates access and refresh token
func CreateTokens(userID string) (map[string]string, error) {
	tokens := map[string]string{}
	accessToken, err := CreateAccessToken(userID)
	if err != nil {
		return tokens, err
	}
	refreshToken, err := CreateRefreshToken(userID)
	if err != nil {
		return tokens, err
	}
	tokens["access"] = accessToken
	tokens["refresh"] = refreshToken
	return tokens, nil
}

//ExtractToken extracts the jwt token string from the request header
func ExtractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return "", errors.New("Missing authentication token")
	}
	tokenSlice := strings.Split(bearerToken, " ")
	if len(tokenSlice) != 2 || strings.ToLower(tokenSlice[0]) != "bearer" {
		return "", errors.New("Token does not use the bearer scheme")
	}
	return tokenSlice[1], nil
}

//ValidateToken validates the jwt authorization token
func ValidateToken(r *http.Request) error {
	tokenString, err := ExtractToken(r)
	if err != nil {
		return err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return errors.New("invalid token claims")
	}
	tokenType := token.Claims.(jwt.MapClaims)["token_type"]
	if tokenType != "access" {
		return errors.New("Token must be an access token")
	}
	return nil
}

//GetIDFromRefreshToken validates the jwt refresh token and extracts the user id
func GetIDFromRefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return "", errors.New("invalid token claims")
	}
	tokenType := token.Claims.(jwt.MapClaims)["token_type"]
	if tokenType != "refresh" {
		return "", errors.New("Token must be a refresh token")
	}
	userID := token.Claims.(jwt.MapClaims)["user_id"]
	if _, ok := userID.(string); !ok {
		return "", errors.New("The user id in the token is not a uuid string")
	}
	return userID.(string), nil
}

//GetIDFromAccessToken validates the jwt access token and extracts the user id
func GetIDFromAccessToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return "", errors.New("invalid token claims")
	}
	tokenType := token.Claims.(jwt.MapClaims)["token_type"]
	if tokenType != "access" {
		return "", errors.New("Token must be an access token")
	}
	userID := token.Claims.(jwt.MapClaims)["user_id"]
	if _, ok := userID.(string); !ok {
		return "", errors.New("The user id in the token is not a uuid string")
	}
	return userID.(string), nil
}
