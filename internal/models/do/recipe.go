package do

import "time"

type Recipe struct {
	ID           uint64        `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Cuisine      string        `json:"cuisine"`
	PhotoURL     string        `json:"photo_url"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at,omitempty"`
	Ingredients  []Ingredient  `json:"ingredients" gorm:"many2many:recipe_ingredients;"`
	Instructions []Instruction `json:"instructions,omitempty" gorm:"foreignKey:RecipeID"`
}

type RecipeIngredient struct {
	ID           uint64     `json:"id"`
	RecipeID     uint64     `json:"recipe_id"`
	IngredientID uint64     `json:"ingredient_id"`
	Quantity     string     `json:"quantity"`
	Unit         string     `json:"unit"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
