// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UpdateLessonUseCase handles lesson updates.
type UpdateLessonUseCase struct{}

func NewUpdateLessonUseCase() *UpdateLessonUseCase { return &UpdateLessonUseCase{} }

func (uc *UpdateLessonUseCase) Execute(ctx context.Context, lessonID string, req *dto.UpdateLessonRequest) (*dto.LessonResponse, error) {
	return nil, nil
}
