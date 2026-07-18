// Package progress contains progress tracking use cases.
package progress

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// GetCourseProgressUseCase handles getting course completion stats.
type GetCourseProgressUseCase struct {
	progressRepo repository.ProgressRepository
}

// NewGetCourseProgressUseCase creates a new GetCourseProgressUseCase.
func NewGetCourseProgressUseCase(progressRepo repository.ProgressRepository) *GetCourseProgressUseCase {
	return &GetCourseProgressUseCase{progressRepo: progressRepo}
}

// Execute returns the user's completion stats for a course.
func (uc *GetCourseProgressUseCase) Execute(ctx context.Context, userID, courseID string) (*dto.CourseProgressResponse, error) {
	completed, total, err := uc.progressRepo.GetCourseCompletionStats(ctx, userID, courseID)
	if err != nil {
		return nil, err
	}
	pct := 0.0
	if total > 0 {
		pct = float64(completed) / float64(total) * 100
	}
	return &dto.CourseProgressResponse{
		CourseID:         courseID,
		CompletedLessons: completed,
		TotalLessons:     total,
		ProgressPct:      pct,
	}, nil
}
