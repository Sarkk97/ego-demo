package repositories

import (
	"net/http"

	"github.com/jenlesamuel/recipe-collection/connection"
	"github.com/jenlesamuel/recipe-collection/httperrors"
	"github.com/jenlesamuel/recipe-collection/models"
	"github.com/jinzhu/gorm"
)

type (
	// RecipeRepository is a Recipe collection manager
	RecipeRepository interface {
		AddRecipe(recipe *models.Recipe) error
	}

	recipeDBRepository struct {
		db *gorm.DB
	}
)

var db *gorm.DB

func init() {
	db = connection.GetDB()
	db.AutoMigrate(&models.Recipe{})
	db.AutoMigrate(&models.RecipeIngredient{})
}

//NewRecipeRepository returns an instance of a RecipeRepository impelmentation
func NewRecipeRepository() RecipeRepository {
	return &recipeDBRepository{
		db: db,
	}
}

func (recipeRepository recipeDBRepository) AddRecipe(recipe *models.Recipe) error {
	if err := recipeRepository.db.Create(recipe).Error; err != nil {
		return httperrors.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: httperrors.ServerError,
		}
	}

	return nil
}
