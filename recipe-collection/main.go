package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jenlesamuel/recipe-collection/handlers"
)

func registerRouteHandlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/login", handlers.Login).Methods(http.MethodPost)

	router.HandleFunc("/recipe", handlers.CreateRecipe).Methods(http.MethodPost)

	return router
}

func main() {
	router := registerRouteHandlers()

	log.Fatalln(http.ListenAndServe(":8080", router))
}

/*func structToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
*/
