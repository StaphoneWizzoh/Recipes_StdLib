package main

import (
	"Recipes_StdLib/recipes"
	"Recipes_StdLib/views"
	"log"
	"net/http"
)

func main() {

	store, err := recipes.NewGormStore()
	if err != nil{
		log.Println("Internal database error: ", err)
	}
	recipesHandler := views.NewRecipeHandler(store)

	mux := http.NewServeMux()

	mux.Handle("/", &views.HomeHandler{})
	mux.Handle("/recipes", recipesHandler)
	mux.Handle("/recipes/", recipesHandler)

	log.Println("Server running on port: 8000")
	http.ListenAndServe(":8000", mux)
}
