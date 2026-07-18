// Package progress contains progress tracking use cases.
package progress

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ListProgressUseCase lists all lesson progress for a user.
type ListProgressUseCase struct {
	progressRepo repository.ProgressRepository
}

// NewListProgressUseCase creates a new ListProgressUseCase.
func NewListProgressUseCase(progressRepo repository.ProgressRepository) *ListProgressUseCase {
	return &ListProgressUseCase{progressRepo: progressRepo}
}

// Execute returns every progress record for the user (module_id included).
func (uc *ListProgressUseCase) Execute(ctx context.Context, userID string) ([]*dto.ProgressResponse, error) {
	list, err := uc.progressRepo.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	responses := make([]*dto.ProgressResponse, 0, len(list))
	for _, p := range list {
		responses = append(responses, toProgressResponse(p))
	}
	return responses, nil
}
