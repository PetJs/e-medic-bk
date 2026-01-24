// Package module contains module management use cases.
package module

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UpdateModuleUseCase handles module updates.
type UpdateModuleUseCase struct{}

// NewUpdateModuleUseCase creates a new UpdateModuleUseCase.
func NewUpdateModuleUseCase() *UpdateModuleUseCase { return &UpdateModuleUseCase{} }

// Execute updates an existing module.
func (uc *UpdateModuleUseCase) Execute(ctx context.Context, moduleID string, req *dto.UpdateModuleRequest) (*dto.ModuleResponse, error) {
	return nil, nil
}
