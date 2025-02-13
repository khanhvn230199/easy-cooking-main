package router

import (
	"easy-cooking/internal/handler"
	"github.com/gin-gonic/gin"
)

func RecipeRouter(r *gin.Engine, handler *handler.Handler) {
	recipeGroup := r.Group("v0/recipes")
	{
		recipeGroup.GET("", handler.GetRecipes)
		recipeGroup.GET("/:id", handler.GetRecipe)
		recipeGroup.POST("/search", handler.SearchRecipes)
	}
}
