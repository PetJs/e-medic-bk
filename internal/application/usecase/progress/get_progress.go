// Package progress contains progress tracking use cases.
package progress

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// GetProgressUseCase handles getting user progress.
type GetProgressUseCase struct{}

func NewGetProgressUseCase() *GetProgressUseCase { return &GetProgressUseCase{} }

func (uc *GetProgressUseCase) Execute(ctx context.Context, userID, lessonID string) (*dto.ProgressResponse, error) {
	return nil, nil
}
