package recipes

import (
	"gorm.io/gorm"
)



type Recipe struct {
	gorm.Model
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	gorm.Model
	Name string `json:"name"`
}

type RecipeStore interface {
	Add(name string, recipe Recipe) error
	Get(name string) (Recipe, error)
	Update(name string, recipe Recipe) error
	List() (map[string]Recipe, error)
	Remove(name string) error
}

// func InitDatabase(){
// 	// Connect to SQLite database

// 	var err error
// 	db, err = gorm.Open(sqlite.Open("../recipe.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to database")
// 	}

// 	// Auto Migrate model structs
// 	db.AutoMigrate(&User{}, &Message{}) 
// }