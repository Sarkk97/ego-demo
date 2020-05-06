package tests

import (
	"ego/user/models"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/go-playground/assert.v1"
)

func TestRepoCreateUser(t *testing.T) {
	var err error
	if err = refreshUserTable(); err != nil {
		t.Errorf("Can't refresh table")
	}

	user := models.User{
		Phone: "08023963212",
		Email: "testemail@gmail.com",
		PIN:   "1997",
	}

	if err = repo.CreateUser(&user); err != nil {
		t.Errorf("Error with repo: %v", err)
	}
}

func TestRepoGetUsers(t *testing.T) {
	var err error
	if err = refreshUserTable(); err != nil {
		t.Errorf("Can't refresh table")
	}
	_ = repo.CreateUser(&models.User{
		Phone: "08023963212",
		Email: "testemail@gmail.com",
		PIN:   "1997",
	})

	users, err := repo.GetAllUsers()
	if err != nil {
		t.Errorf("Error with repo: %v", err)
	}
	assert.Equal(t, len(users), 1)

}
