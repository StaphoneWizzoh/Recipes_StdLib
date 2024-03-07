package recipes

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	NotFoundError = errors.New("Not Found")
)

type GormStore struct {
	db *gorm.DB
}



func NewGormStore() (*GormStore, error) {
	db, err := gorm.Open(sqlite.Open("../recipe.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Recipe{}, &Ingredient{}); err != nil {
		return nil, err
	}

	return &GormStore{db: db}, nil
}

func (s *GormStore) Add(name string, recipe Recipe) error {
    recipe.Name = name
    if err := s.db.Create(&recipe).Error; err != nil {
        return err
    }
    return nil
}

func (s *GormStore) Get(name string) (Recipe, error) {
	var recipe Recipe
	result := s.db.Where("name = ?", name).First(&recipe)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Recipe{}, gorm.ErrRecordNotFound
		}
		return Recipe{}, result.Error
	}
	return recipe, nil
}

func (s *GormStore) List() ([]Recipe, error) {
	var recipes []Recipe
	if err := s.db.Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func (s *GormStore) Update(name string, recipe Recipe) error {
	result := s.db.Model(&Recipe{}).Where("name = ?", name).Updates(recipe)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}

func (s *GormStore) Remove(name string) error {
	result := s.db.Where("name = ?", name).Delete(&Recipe{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}