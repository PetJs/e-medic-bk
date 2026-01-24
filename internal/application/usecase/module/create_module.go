// Package module contains module management use cases.
package module

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// CreateModuleUseCase handles module creation.
type CreateModuleUseCase struct{}

// NewCreateModuleUseCase creates a new CreateModuleUseCase.
func NewCreateModuleUseCase() *CreateModuleUseCase { return &CreateModuleUseCase{} }

// Execute creates a new module.
func (uc *CreateModuleUseCase) Execute(ctx context.Context, req *dto.CreateModuleRequest) (*dto.ModuleResponse, error) {
	return nil, nil
}
