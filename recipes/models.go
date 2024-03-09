package recipes

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)



type Recipe struct {
	gorm.Model
	ID			uuid.UUID	 `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	gorm.Model
	ID			uuid.UUID	 `json:"id"`
	Name string `json:"name"`
}

type RecipeStore interface {
	Add(name string, recipe Recipe) error
	Get(name string) (Recipe, error)
	Update(name string, recipe Recipe) error
	List() (map[string]Recipe, error)
	Remove(name string) error
}
