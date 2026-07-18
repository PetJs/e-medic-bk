// Package progress contains progress tracking use cases.
package progress

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// GetProgressUseCase handles getting a user's progress on one lesson.
type GetProgressUseCase struct {
	progressRepo repository.ProgressRepository
}

// NewGetProgressUseCase creates a new GetProgressUseCase.
func NewGetProgressUseCase(progressRepo repository.ProgressRepository) *GetProgressUseCase {
	return &GetProgressUseCase{progressRepo: progressRepo}
}

// Execute returns the user's progress on a lesson; a zero-progress response
// is returned when the lesson has not been started.
func (uc *GetProgressUseCase) Execute(ctx context.Context, userID, lessonID string) (*dto.ProgressResponse, error) {
	p, err := uc.progressRepo.GetByUserAndLesson(ctx, userID, lessonID)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return &dto.ProgressResponse{UserID: userID, LessonID: lessonID}, nil
	}
	return toProgressResponse(p), nil
}
