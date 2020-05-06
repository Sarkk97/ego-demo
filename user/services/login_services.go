package services

import (
	"ego/user/auth"
	"ego/user/models"
	"ego/user/repositories"
)

//LoginService is a struct for login services
type LoginService struct {
	Repo *repositories.GormRepo
}

//NewLoginService is constructor for LoginService
func NewLoginService(r *repositories.GormRepo) *LoginService {
	return &LoginService{
		Repo: r,
	}
}

//LoginUser is a login service method to login a user
func (l *LoginService) LoginUser(user models.LoginUser) (map[string]string, error) {
	loggedInUser, err := l.Repo.LoginUser(user)
	tokens := map[string]string{}
	if err != nil {
		return tokens, err
	}
	//create user tokens
	tokens, err = auth.CreateTokens(loggedInUser.ID)
	if err != nil {
		return tokens, err
	}
	//update user last_login
	_ = l.Repo.UpdateUserLogin(loggedInUser)
	return tokens, nil
}

//RefreshToken is a service that generates a new access token from a valid refresh token
func (l *LoginService) RefreshToken(token models.RefreshToken) (map[string]string, error) {
	payload := map[string]string{}
	userID, err := auth.GetIDFromRefreshToken(token.Refresh)
	if err != nil {
		return payload, err
	}
	accessToken, err := auth.CreateAccessToken(userID)
	if err != nil {
		return payload, err
	}
	payload["access"] = accessToken
	return payload, nil

}
