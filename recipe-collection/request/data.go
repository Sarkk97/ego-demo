package request

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/jenlesamuel/recipe-collection/models"
)

type (
	//RecipeData represents a recipe request body
	RecipeData struct {
		Name        string            `json:"name"`
		PrepTime    uint              `json:"prepTime"`
		Difficulty  uint              `json:"difficulty"`
		Ingredients []*ingredientData `json:"ingredients"`
	}

	ingredientData struct {
		RecipeID string `json:"recipeId"`
		Name     string `json:"name"`
		UOM      string `json:"uom"`
		Quantity uint   `json:"quantity"`
		ImageURL string `json:"imageURL"`
	}
)

//Validate validates RecipeData fields
func (recipeData *RecipeData) Validate() url.Values {
	return make(url.Values)
}

//ToRecipe transforms recipe request data to Recipe model
func (recipeData RecipeData) ToRecipe() *models.Recipe {
	recipeIngredients := []*models.RecipeIngredient{}

	for _, ingredientData := range recipeData.Ingredients {
		recipeIngredient := &models.RecipeIngredient{
			ID:        uuid.New().String(),
			RecipeID:  ingredientData.RecipeID,
			Name:      ingredientData.Name,
			UOM:       ingredientData.UOM,
			Quantity:  ingredientData.Quantity,
			ImageURL:  ingredientData.ImageURL,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		recipeIngredients = append(recipeIngredients, recipeIngredient)
	}

	return &models.Recipe{
		ID:          uuid.New().String(),
		Name:        recipeData.Name,
		PrepTime:    recipeData.PrepTime,
		Difficulty:  recipeData.Difficulty,
		Ingredients: recipeIngredients,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}
