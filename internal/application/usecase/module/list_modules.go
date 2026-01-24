// Package module contains module management use cases.
package module

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ListModulesUseCase handles listing modules by course.
type ListModulesUseCase struct{}

// NewListModulesUseCase creates a new ListModulesUseCase.
func NewListModulesUseCase() *ListModulesUseCase { return &ListModulesUseCase{} }

// Execute lists modules for a course.
func (uc *ListModulesUseCase) Execute(ctx context.Context, courseID string) ([]*dto.ModuleResponse, error) {
	return nil, nil
}
