package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jenlesamuel/recipe-collection/httperrors"

	"github.com/jenlesamuel/recipe-collection/factories"
	"github.com/jenlesamuel/recipe-collection/request"
	"github.com/jenlesamuel/recipe-collection/respond"
)

//CreateRecipe handles request to create a recipe
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	recipeData := &request.RecipeData{}

	json.NewDecoder(r.Body).Decode(recipeData)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	//TODO: Implement validation in RecipeData
	if errBag := recipeData.Validate(); len(errBag) != 0 {
		respond.WithError(
			w,
			errBag,
			http.StatusBadRequest,
			headers,
		)

		return
	}

	recipe, err := factories.CreateRecipeService().CreateRecipe(recipeData)

	if err != nil {
		respond.WithError(
			w,
			err.Error(),
			err.(httperrors.HTTPError).Status,
			headers,
		)

		return
	}

	respond.WithSuccess(
		w,
		recipe,
		http.StatusCreated,
		headers,
	)

	return
}
