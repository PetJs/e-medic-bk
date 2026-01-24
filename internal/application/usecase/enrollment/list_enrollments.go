// Package enrollment contains enrollment use cases.
package enrollment

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ListEnrollmentsUseCase handles listing user enrollments.
type ListEnrollmentsUseCase struct{}

func NewListEnrollmentsUseCase() *ListEnrollmentsUseCase { return &ListEnrollmentsUseCase{} }

func (uc *ListEnrollmentsUseCase) Execute(ctx context.Context, userID string) ([]*dto.EnrollmentResponse, error) {
	return nil, nil
}
