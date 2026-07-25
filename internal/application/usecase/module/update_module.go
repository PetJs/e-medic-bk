// Package module contains module management use cases.
package module

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// UpdateModuleUseCase handles module updates.
type UpdateModuleUseCase struct {
	moduleRepo repository.ModuleRepository
	storageSvc service.StorageService
	coverGen   *ModuleCoverGenerator
}

// NewUpdateModuleUseCase creates a new UpdateModuleUseCase.
func NewUpdateModuleUseCase(
	moduleRepo repository.ModuleRepository,
	storageSvc service.StorageService,
	coverGen *ModuleCoverGenerator,
) *UpdateModuleUseCase {
	return &UpdateModuleUseCase{moduleRepo: moduleRepo, storageSvc: storageSvc, coverGen: coverGen}
}

// Execute updates a module.
func (uc *UpdateModuleUseCase) Execute(ctx context.Context, id string, req *dto.UpdateModuleRequest) (*dto.ModuleResponse, error) {
	module, err := uc.moduleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if module == nil {
		return nil, ErrModuleNotFound
	}

	contentChanged := (req.Title != nil && *req.Title != module.Title) ||
		(req.Description != nil && *req.Description != module.Description)

	if req.Title != nil {
		module.Title = *req.Title
	}
	if req.Description != nil {
		module.Description = *req.Description
	}
	if req.Order != nil {
		module.Order = *req.Order
	}
	if req.IsPremium != nil {
		module.IsPremium = *req.IsPremium
	}

	if err := uc.moduleRepo.Update(ctx, module); err != nil {
		return nil, err
	}
	if contentChanged {
		uc.coverGen.Trigger(module.ID, module.Title, module.Description)
	}
	return ToModuleResponse(ctx, module, uc.storageSvc), nil
}
