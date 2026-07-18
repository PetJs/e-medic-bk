// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// UpdateLessonUseCase handles lesson updates.
type UpdateLessonUseCase struct {
	lessonRepo repository.LessonRepository
}

// NewUpdateLessonUseCase creates a new UpdateLessonUseCase.
func NewUpdateLessonUseCase(lessonRepo repository.LessonRepository) *UpdateLessonUseCase {
	return &UpdateLessonUseCase{lessonRepo: lessonRepo}
}

// Execute updates a lesson.
func (uc *UpdateLessonUseCase) Execute(ctx context.Context, id string, req *dto.UpdateLessonRequest) (*dto.LessonResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	if req.Title != nil {
		lesson.Title = *req.Title
	}
	if req.Description != nil {
		lesson.Description = *req.Description
	}
	if req.Order != nil {
		lesson.Order = *req.Order
	}
	if req.Duration != nil {
		lesson.Duration = *req.Duration
	}

	if err := uc.lessonRepo.Update(ctx, lesson); err != nil {
		return nil, err
	}
	return toLessonResponse(lesson), nil
}
