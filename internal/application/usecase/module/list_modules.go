// Package module contains module management use cases.
package module

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ListModulesUseCase handles listing modules.
type ListModulesUseCase struct {
	moduleRepo repository.ModuleRepository
}

// NewListModulesUseCase creates a new ListModulesUseCase.
func NewListModulesUseCase(moduleRepo repository.ModuleRepository) *ListModulesUseCase {
	return &ListModulesUseCase{moduleRepo: moduleRepo}
}

// Execute lists modules — by course when courseID is set, otherwise across all published courses.
func (uc *ListModulesUseCase) Execute(ctx context.Context, courseID string) ([]*dto.ModuleResponse, error) {
	var (
		modules []*entity.Module
		err     error
	)
	if courseID != "" {
		modules, err = uc.moduleRepo.ListByCourse(ctx, courseID)
	} else {
		modules, err = uc.moduleRepo.ListAll(ctx)
	}
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.ModuleResponse, 0, len(modules))
	for _, m := range modules {
		responses = append(responses, toModuleResponse(m))
	}
	return responses, nil
}

// GetModuleUseCase handles fetching a single module.
type GetModuleUseCase struct {
	moduleRepo repository.ModuleRepository
}

// NewGetModuleUseCase creates a new GetModuleUseCase.
func NewGetModuleUseCase(moduleRepo repository.ModuleRepository) *GetModuleUseCase {
	return &GetModuleUseCase{moduleRepo: moduleRepo}
}

// Execute fetches a module by ID.
func (uc *GetModuleUseCase) Execute(ctx context.Context, id string) (*dto.ModuleResponse, error) {
	module, err := uc.moduleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if module == nil {
		return nil, ErrModuleNotFound
	}
	return toModuleResponse(module), nil
}
