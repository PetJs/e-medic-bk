// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// DeleteLessonUseCase handles lesson deletion.
type DeleteLessonUseCase struct {
	lessonRepo  repository.LessonRepository
	contentRepo repository.ContentRepository
	storage     service.StorageService
}

// NewDeleteLessonUseCase creates a new DeleteLessonUseCase.
func NewDeleteLessonUseCase(
	lessonRepo repository.LessonRepository,
	contentRepo repository.ContentRepository,
	storage service.StorageService,
) *DeleteLessonUseCase {
	return &DeleteLessonUseCase{lessonRepo: lessonRepo, contentRepo: contentRepo, storage: storage}
}

// Execute deletes a lesson, its content rows (via DB cascade), and their
// stored objects. Object cleanup is best effort; the DB is the source of truth.
func (uc *DeleteLessonUseCase) Execute(ctx context.Context, id string) error {
	lesson, err := uc.lessonRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if lesson == nil {
		return ErrLessonNotFound
	}

	if contents, err := uc.contentRepo.ListByLesson(ctx, id); err == nil {
		for _, content := range contents {
			_ = uc.storage.Delete(ctx, content.StorageKey)
		}
	}
	return uc.lessonRepo.Delete(ctx, id)
}
