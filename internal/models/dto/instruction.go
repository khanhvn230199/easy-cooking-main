package dto

import "easy-cooking/internal/models/do"

type InstructionResponse struct {
	ID          uint64 `json:"id"`
	Step        int    `json:"step"`
	Description string `json:"description"`
}

func ToInstructionResponse(inst do.Instruction) InstructionResponse {
	return InstructionResponse{
		ID:          inst.ID,
		Step:        inst.StepNumber,
		Description: inst.Description,
	}
}
