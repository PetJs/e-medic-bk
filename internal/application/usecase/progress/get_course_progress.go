// Package progress contains progress tracking use cases.
package progress

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// GetCourseProgressUseCase handles getting course completion stats.
type GetCourseProgressUseCase struct{}

func NewGetCourseProgressUseCase() *GetCourseProgressUseCase { return &GetCourseProgressUseCase{} }

func (uc *GetCourseProgressUseCase) Execute(ctx context.Context, userID, courseID string) (*dto.CourseProgressResponse, error) {
	return nil, nil
}
