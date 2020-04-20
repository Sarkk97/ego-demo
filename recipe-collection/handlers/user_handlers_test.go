package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/joho/godotenv/autoload" //Autoload env variables
)

func TestCreateUserReturnsValidationErrorGivenBadRequest(t *testing.T) {
	jsonStr := []byte(`{
		"name": "Davo04",
		"email": "davo04@gmail.com",
		"password": "helives"
	}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("Did not return 400 - Bad Request for invalid request body")
	}
}
