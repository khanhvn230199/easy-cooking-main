package handler

import (
	"easy-cooking/internal/service"
	"time"

	"gorm.io/gorm"
)

type Handler struct {
	DB            *gorm.DB
	recipeService service.RecipeService
}

func NewHandler(db *gorm.DB, timeout time.Duration) *Handler {
	return &Handler{
		DB:            db,
		recipeService: service.NewRecipeService(db, timeout),
	}
}
