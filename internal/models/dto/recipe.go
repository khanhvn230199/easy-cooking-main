package dto

import (
	"easy-cooking/internal/models/do"
	"time"
)

type RecipeResponse struct {
	ID           uint64                `json:"id"`
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	Cuisine      string                `json:"cuisine"`
	PhotoURL     string                `json:"photo_url"`
	CreatedAt    time.Time             `json:"created_at"`
	Ingredients  []IngredientResponse  `json:"ingredients"`
	Instructions []InstructionResponse `json:"instructions,omitempty"`
}

func ToRecipeResponse(recipe do.Recipe) RecipeResponse {
	ingredientResponses := make([]IngredientResponse, len(recipe.Ingredients))
	for i, ing := range recipe.Ingredients {
		ingredientResponses[i] = ToIngredientResponse(ing)
	}

	instructionResponses := make([]InstructionResponse, len(recipe.Instructions))
	for i, inst := range recipe.Instructions {
		instructionResponses[i] = ToInstructionResponse(inst)
	}

	return RecipeResponse{
		ID:           recipe.ID,
		Title:        recipe.Title,
		Description:  recipe.Description,
		Cuisine:      recipe.Cuisine,
		PhotoURL:     recipe.PhotoURL,
		CreatedAt:    recipe.CreatedAt,
		Ingredients:  ingredientResponses,
		Instructions: instructionResponses,
	}
}

func ToRecipeResponses(recipes []*do.Recipe) []RecipeResponse {
	responses := make([]RecipeResponse, len(recipes))
	for i, recipe := range recipes {
		if recipe == nil {
			continue
		}
		responses[i] = ToRecipeResponse(*recipe)
	}
	return responses
}

type RecipeSearchRequest struct {
	Keyword     string   `form:"keyword" json:"keyword"`
	Ingredients []string `form:"ingredients" json:"ingredients"`
	Cuisine     string   `form:"cuisine" json:"cuisine"`

	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`

	SortBy    string `form:"sort_by" json:"sort_by"`
	SortOrder string `form:"sort_order" json:"sort_order" binding:"omitempty,oneof=asc desc"`

	MinPreparationTime *int       `form:"min_prep_time" json:"min_prep_time"`
	MaxPreparationTime *int       `form:"max_prep_time" json:"max_prep_time"`
	CreatedAfter       *time.Time `form:"created_after" json:"created_after"`
}

// Validate method (optional)
func (r *RecipeSearchRequest) Validate() error {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.PageSize < 1 {
		r.PageSize = 10
	}

	if r.PageSize > 100 {
		r.PageSize = 100
	}

	return nil
}

type SearchRecipeResponse struct {
	Recipes    []RecipeResponse `json:"recipes"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalPages int              `json:"total_pages"`
}

// SearchRecipeResponse
func NewSearchRecipeResponse(recipes []*do.Recipe, total int64, page int, pageSize int) SearchRecipeResponse {
	totalPages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPages++
	}

	return SearchRecipeResponse{
		Recipes:    ToRecipeResponses(recipes),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}
