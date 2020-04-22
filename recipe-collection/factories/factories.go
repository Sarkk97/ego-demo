package factories

import (
	"github.com/jenlesamuel/recipe-collection/repositories"
	"github.com/jenlesamuel/recipe-collection/services"
)

//CreateRecipeService makes a new RecipeService instance
func CreateRecipeService() *services.RecipeService {
	return services.NewRecipeService(repositories.NewRecipeRepository())
}
