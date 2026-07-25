// Package module contains module management use cases.
package module

import (
	"context"

	"emedic-bk/internal/domain/repository"
)

// RegenerateCoverUseCase re-triggers cover image generation for a module —
// used to backfill modules created before this feature and to retry failures.
type RegenerateCoverUseCase struct {
	moduleRepo repository.ModuleRepository
	coverGen   *ModuleCoverGenerator
}

// NewRegenerateCoverUseCase creates a new RegenerateCoverUseCase.
func NewRegenerateCoverUseCase(moduleRepo repository.ModuleRepository, coverGen *ModuleCoverGenerator) *RegenerateCoverUseCase {
	return &RegenerateCoverUseCase{moduleRepo: moduleRepo, coverGen: coverGen}
}

// Execute resets the module's cover status to pending and re-triggers generation.
func (uc *RegenerateCoverUseCase) Execute(ctx context.Context, id string) error {
	module, err := uc.moduleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if module == nil {
		return ErrModuleNotFound
	}
	if err := uc.moduleRepo.UpdateCoverImage(ctx, id, "", "pending"); err != nil {
		return err
	}
	uc.coverGen.Trigger(module.ID, module.Title, module.Description)
	return nil
}
