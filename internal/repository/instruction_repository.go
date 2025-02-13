package repository

import (
	"context"
	"easy-cooking/internal/models/do"
	"gorm.io/gorm"
)

type GetInstructionsFilter struct {
	recipeIDs []int64
}
type instructionRepository struct {
	db *gorm.DB
}

type InstructionRepository interface {
	GetInstructions(ctx context.Context, filter GetInstructionsFilter) ([]*do.Instruction, error)
}

var _ InstructionRepository = (*instructionRepository)(nil)

func NewInstructionRepository(db *gorm.DB) *instructionRepository {
	return &instructionRepository{db: db}
}

func (r *instructionRepository) GetInstructions(ctx context.Context, filter GetInstructionsFilter) ([]*do.Instruction, error) {
	var instructions []*do.Instruction
	query := r.db.WithContext(ctx).Model(&do.Instruction{})
	if len(filter.recipeIDs) > 0 {
		query = query.Where("recipe_id IN ?", filter.recipeIDs)
	}
	if err := r.db.WithContext(ctx).Find(&instructions).Error; err != nil {
		return nil, err
	}
	return instructions, nil
}
