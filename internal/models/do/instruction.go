package do

import "time"

type Instruction struct {
	ID          uint64     `json:"id"`
	RecipeID    uint64     `json:"recipe_id"`
	StepNumber  int        `json:"step_number"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
