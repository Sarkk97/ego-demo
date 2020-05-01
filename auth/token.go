package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateAccessToken generates the JWT access token with claims
func CreateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["token_type"] = "access"
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

//CreateRefreshToken generates the JWT refresh token with claims
func CreateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["token_type"] = "refresh"
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hours
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
