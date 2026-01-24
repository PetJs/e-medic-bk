// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// GetLessonUseCase handles getting lesson details with content.
type GetLessonUseCase struct{}

func NewGetLessonUseCase() *GetLessonUseCase { return &GetLessonUseCase{} }

func (uc *GetLessonUseCase) Execute(ctx context.Context, lessonID, userID string) (*dto.LessonDetailResponse, error) {
	return nil, nil
}
