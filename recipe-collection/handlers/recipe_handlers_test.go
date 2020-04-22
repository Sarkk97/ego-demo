package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRecipeSuccess(t *testing.T) {
	data := []byte(`{
		"name": "Salmon Limone",
		"prepTime": 30,
		"difficulty": 1,
		"ingredients": [
			{
				"name": "scallions", "uom": "unit", "quantity": 2, 
				"imageURL": "https://img.recipe.com/554a301f4dab71626c8b4569-58c7e527.png"
			},
			{
				"name": "Israeli Couscous", "uom": "cup", "quantity": 0.5, 
				"imageURL":"https://img.recipe.com/554a301f4dab71626c8b4569-58c7e527.png"
			}
		]
	}`)

	req, _ := http.NewRequest("POST", "/recipe", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateRecipe)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Error("Did not return 201 - created for valid request")
	}
}
