// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// CreateLessonUseCase handles lesson creation.
type CreateLessonUseCase struct{}

func NewCreateLessonUseCase() *CreateLessonUseCase { return &CreateLessonUseCase{} }

func (uc *CreateLessonUseCase) Execute(ctx context.Context, req *dto.CreateLessonRequest) (*dto.LessonResponse, error) {
	return nil, nil
}
