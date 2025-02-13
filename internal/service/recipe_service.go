package service

import (
	"context"
	"easy-cooking/internal/models/dto"
	"easy-cooking/internal/repository"
	"gorm.io/gorm"
	"time"
)

type RecipeService interface {
	GetRecipes(ctx context.Context, params dto.RecipeSearchRequest) (*dto.SearchRecipeResponse, error)
	SearchRecipes(ctx context.Context, params dto.RecipeSearchRequest) (*dto.SearchRecipeResponse, error)
	GetRecipeByID(ctx context.Context, id int64) (*dto.RecipeResponse, error)
}

type recipeService struct {
	timeout          time.Duration
	recipeRepository repository.RecipeRepository
}

func NewRecipeService(db *gorm.DB, timeout time.Duration) *recipeService {
	return &recipeService{
		timeout:          timeout,
		recipeRepository: repository.NewRecipeRepository(db),
	}
}

var _ RecipeService = (*recipeService)(nil)

func (s *recipeService) GetRecipes(ctx context.Context, params dto.RecipeSearchRequest) (*dto.SearchRecipeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	searchCriteria := repository.SearchCriteria{
		Page:      params.Page,
		PageSize:  params.PageSize,
		SortBy:    params.SortBy,
		SortOrder: params.SortOrder,
	}
	recipes, total, err := s.recipeRepository.GetRecipes(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}
	searchResponse := dto.NewSearchRecipeResponse(
		recipes,
		total,
		params.Page,
		params.PageSize,
	)
	return &searchResponse, nil
}

func (s *recipeService) SearchRecipes(ctx context.Context, params dto.RecipeSearchRequest) (*dto.SearchRecipeResponse, error) {
	// Validate request
	if err := params.Validate(); err != nil {
		return nil, err
	}

	searchCriteria := repository.SearchCriteria{
		Keyword:     params.Keyword,
		Ingredients: params.Ingredients,
		Cuisine:     params.Cuisine,
		Page:        params.Page,
		PageSize:    params.PageSize,
		SortBy:      params.SortBy,
		SortOrder:   params.SortOrder,
	}

	recipes, total, err := s.recipeRepository.SearchRecipes(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	searchResponse := dto.NewSearchRecipeResponse(
		recipes,
		total,
		params.Page,
		params.PageSize,
	)

	return &searchResponse, nil
}

func (s *recipeService) GetRecipeByID(ctx context.Context, id int64) (*dto.RecipeResponse, error) {

	recipe, err := s.recipeRepository.GetRecipeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := dto.ToRecipeResponse(recipe)
	return &resp, nil
}
