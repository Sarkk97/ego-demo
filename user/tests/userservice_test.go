package tests

import (
	"ego/user/models"
	"testing"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestUserCreateService(t *testing.T) {
	var err error
	if err = refreshUserTable(); err != nil {
		t.Errorf("Error refreshing table: %v", err)
	}

	user := models.User{
		ID:    uuid.New().String(),
		Phone: "08023963212",
		Email: "testemail@gmail.com",
		PIN:   "1997",
	}

	err = service.CreateUser(&user)
	if err != nil {
		t.Errorf("Error with service: %v", err)
	}

	// assert.Equal(t, user.ID, uint32(1))
	if err = user.ConfirmPIN("1997"); err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}
}
