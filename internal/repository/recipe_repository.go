package repository

import (
	"context"
	"easy-cooking/internal/models/do"
	"gorm.io/gorm"
)

type recipeRepository struct {
	db *gorm.DB
}

type RecipeRepository interface {
	GetRecipes(ctx context.Context, filter SearchCriteria) ([]*do.Recipe, int64, error)
	SearchRecipes(ctx context.Context, criteria SearchCriteria) ([]*do.Recipe, int64, error)
	GetRecipeByID(ctx context.Context, id int64) (do.Recipe, error)
}

var _ RecipeRepository = (*recipeRepository)(nil)

func NewRecipeRepository(db *gorm.DB) *recipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) GetRecipes(ctx context.Context, filter SearchCriteria) ([]*do.Recipe, int64, error) {
	var recipes []*do.Recipe
	query := r.db.WithContext(ctx).Model(&do.Recipe{})
	query = query.Preload("Ingredients").
		Preload("Instructions", func(db *gorm.DB) *gorm.DB {
			return db.Order("step_number ASC")
		})
	if filter.SortBy != "" {
		orderClause := filter.SortBy
		if filter.SortOrder != "" {
			orderClause += " " + filter.SortOrder
		}
		query = query.Order(orderClause)
	}

	var total int64
	countQuery := query
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filter.Page - 1) * filter.PageSize
	query = query.Offset(offset).Limit(filter.PageSize)
	if err := query.Find(&recipes).Error; err != nil {
		return nil, 0, err
	}
	return recipes, total, nil
}

type SearchCriteria struct {
	Keyword     string
	Ingredients []string
	Cuisine     string
	Page        int
	PageSize    int
	SortBy      string
	SortOrder   string
}

func (r *recipeRepository) SearchRecipes(ctx context.Context, criteria SearchCriteria) ([]*do.Recipe, int64, error) {
	query := r.db.Model(&do.Recipe{})
	query = query.Preload("Ingredients").
		Preload("Instructions", func(db *gorm.DB) *gorm.DB {
			return db.Order("step_number ASC")
		})
	if criteria.Keyword != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?",
			"%"+criteria.Keyword+"%",
			"%"+criteria.Keyword+"%")
	}

	if criteria.Cuisine != "" {
		query = query.Where("cuisine = ?", criteria.Cuisine)
	}

	if len(criteria.Ingredients) > 0 {
		query = query.Joins("JOIN recipe_ingredients ri ON ri.recipe_id = recipes.id").
			Joins("JOIN ingredients i ON i.id = ri.ingredient_id").
			Where("i.name IN ?", criteria.Ingredients).
			Group("recipes.id")
	}

	if criteria.SortBy != "" {
		orderClause := criteria.SortBy
		if criteria.SortOrder != "" {
			orderClause += " " + criteria.SortOrder
		}
		query = query.Order(orderClause)
	}

	var total int64
	countQuery := query
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (criteria.Page - 1) * criteria.PageSize
	query = query.Offset(offset).Limit(criteria.PageSize)

	var recipes []*do.Recipe
	if err := query.Find(&recipes).Error; err != nil {
		return nil, 0, err
	}

	return recipes, total, nil
}

func (r *recipeRepository) GetRecipeByID(ctx context.Context, id int64) (do.Recipe, error) {
	var recipe do.Recipe
	query := r.db.Model(&do.Recipe{})
	query = query.Preload("Ingredients").
		Preload("Instructions", func(db *gorm.DB) *gorm.DB {
			return db.Order("step_number ASC")
		})
	if err := query.Where("id = ?", id).First(&recipe).Error; err != nil {
		return recipe, err
	}
	return recipe, nil
}
