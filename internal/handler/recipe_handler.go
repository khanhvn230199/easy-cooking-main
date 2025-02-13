package handler

import (
	"easy-cooking/internal/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetRecipes(c *gin.Context) {
	page := c.Query("page")
	perPage := c.Query("per_page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	perPageInt, err := strconv.Atoi(perPage)
	if err != nil {
		perPageInt = 10
	}

	searchParams := dto.RecipeSearchRequest{
		Page:      pageInt,
		PageSize:  perPageInt,
		SortBy:    c.Query("sort_by"),
		SortOrder: c.Query("sort_order"),
	}
	recipes, err := h.recipeService.GetRecipes(c.Request.Context(), searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func (h *Handler) GetRecipe(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
	}
	recipes, err := h.recipeService.GetRecipeByID(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func (h *Handler) SearchRecipes(c *gin.Context) {
	var searchParams dto.RecipeSearchRequest

	if err := c.ShouldBind(&searchParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid search parameters",
			"details": err.Error(),
		})
		return
	}

	searchResponse, err := h.recipeService.SearchRecipes(c.Request.Context(), searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch recipes",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, searchResponse)
}
