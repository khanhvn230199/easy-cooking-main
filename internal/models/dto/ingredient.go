package dto

import "easy-cooking/internal/models/do"

type IngredientResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func ToIngredientResponse(ing do.Ingredient) IngredientResponse {
	return IngredientResponse{
		ID:   ing.ID,
		Name: ing.Name,
	}
}
