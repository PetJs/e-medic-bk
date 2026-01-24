// Package progress contains progress tracking use cases.
package progress

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UpdateProgressUseCase handles updating lesson progress.
type UpdateProgressUseCase struct{}

func NewUpdateProgressUseCase() *UpdateProgressUseCase { return &UpdateProgressUseCase{} }

func (uc *UpdateProgressUseCase) Execute(ctx context.Context, userID string, req *dto.UpdateProgressRequest) (*dto.ProgressResponse, error) {
	return nil, nil
}
