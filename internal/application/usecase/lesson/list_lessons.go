// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ListLessonsUseCase handles listing lessons by module.
type ListLessonsUseCase struct{}

func NewListLessonsUseCase() *ListLessonsUseCase { return &ListLessonsUseCase{} }

func (uc *ListLessonsUseCase) Execute(ctx context.Context, moduleID string) ([]*dto.LessonResponse, error) {
	return nil, nil
}
