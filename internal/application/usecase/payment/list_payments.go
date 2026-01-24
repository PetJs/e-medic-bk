// Package payment contains payment use cases.
package payment

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ListPaymentsUseCase handles listing payment history.
type ListPaymentsUseCase struct{}

func NewListPaymentsUseCase() *ListPaymentsUseCase { return &ListPaymentsUseCase{} }

func (uc *ListPaymentsUseCase) Execute(ctx context.Context, userID string, req *dto.ListPaymentsRequest) (*dto.ListPaymentsResponse, error) {
	return nil, nil
}
