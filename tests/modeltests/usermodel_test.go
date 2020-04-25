package modeltests

import (
	"ego/user/models"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateUser(t *testing.T) {
	if err := refreshUserTable(); err != nil {
		t.Errorf("Unable to refresh user table: %v\n", err)
	}
	user := models.User{
		Phone: "090974847847",
		PIN:   "1997",
		Email: "testemail@gmail.com",
	}

	err := db.Create(&user).Error

	if err != nil {
		t.Fatalf("Can not create user data: %v", err)
	}
	assert.Equal(t, user.ID, uint32(1))
	assert.Equal(t, user.PIN, "1997")
	assert.Equal(t, user.Phone, "090974847847")
	assert.Equal(t, user.Email, "testemail@gmail.com")

}

func TestGetAllUsers(t *testing.T) {
	if err := refreshUserTable(); err != nil {
		t.Errorf("Unable to refresh user table: %v\n", err)
	}
	if err := seedUsers(); err != nil {
		t.Errorf("Could not seed users: %v", err)
	}
	users := []models.User{}
	err := db.Find(&users).Error
	if err != nil {
		t.Fatalf("Error runnig query: %v", err)
	}
	assert.Equal(t, len(users), 2)
}

func TestGetAUser(t *testing.T) {
	if err := refreshUserTable(); err != nil {
		t.Errorf("Unable to refresh user table: %v\n", err)
	}
	createdUser, err := seedOneUser()
	// err := seedUsers()
	if err != nil {
		t.Errorf("Could not seed user: %v", err)
	}
	var user []models.User
	err = db.Find(&user).Error
	if err != nil {
		t.Fatalf("Error occured when running query: %v", err)
	}
	assert.Equal(t, len(user), 1)
	assert.Equal(t, user[0].ID, createdUser.ID)
}

func TestUpdateUser(t *testing.T) {
	var err error
	if err = refreshUserTable(); err != nil {
		t.Errorf("Unable to refresh user table: %v\n", err)
	}
	user, err := seedOneUser()
	if err != nil {
		t.Errorf("Could not seed user: %v", err)
	}
	user.Email = "newmail@email.com"
	if err = db.Save(&user).Error; err != nil {
		t.Fatalf("Error running query: %v", err)
	}

	var foundUser models.User
	db.Find(&foundUser)

	assert.Equal(t, foundUser.Email, user.Email)
}
