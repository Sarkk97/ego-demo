package services

import (
	"github.com/jenlesamuel/recipe-collection/models"
	"github.com/jenlesamuel/recipe-collection/repositories"
	"github.com/jenlesamuel/recipe-collection/request"
)

// RecipeService coordinates the processing of a user request
type RecipeService struct {
	recipeRepository repositories.RecipeRepository
}

//CreateRecipe orchestrates recipe creation
func (recipeService RecipeService) CreateRecipe(recipeData *request.RecipeData) (*models.Recipe, error) {
	recipe := recipeData.ToRecipe()

	err := recipeService.recipeRepository.AddRecipe(recipe)

	return recipe, err
}

//NewRecipeService instantiates a RecipeService
func NewRecipeService(recipeRepository repositories.RecipeRepository) *RecipeService {
	return &RecipeService{
		recipeRepository: recipeRepository,
	}
}
