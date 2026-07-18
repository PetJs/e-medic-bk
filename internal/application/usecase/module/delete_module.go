// Package module contains module management use cases.
package module

import (
	"context"

	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// DeleteModuleUseCase handles module deletion.
type DeleteModuleUseCase struct {
	moduleRepo  repository.ModuleRepository
	lessonRepo  repository.LessonRepository
	contentRepo repository.ContentRepository
	storage     service.StorageService
}

// NewDeleteModuleUseCase creates a new DeleteModuleUseCase.
func NewDeleteModuleUseCase(
	moduleRepo repository.ModuleRepository,
	lessonRepo repository.LessonRepository,
	contentRepo repository.ContentRepository,
	storage service.StorageService,
) *DeleteModuleUseCase {
	return &DeleteModuleUseCase{moduleRepo: moduleRepo, lessonRepo: lessonRepo, contentRepo: contentRepo, storage: storage}
}

// Execute deletes a module and everything under it. Lesson and content rows
// go via DB cascade; stored objects are removed best effort first.
func (uc *DeleteModuleUseCase) Execute(ctx context.Context, id string) error {
	module, err := uc.moduleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if module == nil {
		return ErrModuleNotFound
	}

	if lessons, err := uc.lessonRepo.ListByModule(ctx, id); err == nil {
		for _, lesson := range lessons {
			if contents, err := uc.contentRepo.ListByLesson(ctx, lesson.ID); err == nil {
				for _, content := range contents {
					_ = uc.storage.Delete(ctx, content.StorageKey)
				}
			}
		}
	}
	return uc.moduleRepo.Delete(ctx, id)
}
