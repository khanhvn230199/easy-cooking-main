package repository

import (
	"context"
	"easy-cooking/internal/models/do"
	"gorm.io/gorm"
)

type GetIngredientsFilter struct {
	recipeIDs []int64
}
type ingredientRepository struct {
	db *gorm.DB
}

type IngredientRepository interface {
	GetIngredients(ctx context.Context, filter GetIngredientsFilter) ([]*do.Ingredient, error)
}

var _ IngredientRepository = (*ingredientRepository)(nil)

func NewIngredientRepository(db *gorm.DB) *ingredientRepository {
	return &ingredientRepository{db: db}
}

func (r *ingredientRepository) GetIngredients(ctx context.Context, filter GetIngredientsFilter) ([]*do.Ingredient, error) {
	var ingredients []*do.Ingredient
	query := r.db.WithContext(ctx).Model(&do.Ingredient{})
	if len(filter.recipeIDs) > 0 {
		query = query.Joins("JOIN recipe_ingredients ri ON ri.recipe_id = recipes.id").Where("ri.recipe_id IN ?", filter.recipeIDs)
	}
	if err := r.db.WithContext(ctx).Find(&ingredients).Error; err != nil {
		return nil, err
	}
	return ingredients, nil
}
