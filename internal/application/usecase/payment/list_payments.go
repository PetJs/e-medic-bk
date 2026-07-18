// Package payment contains payment use cases.
package payment

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/shared/pagination"
)

// ListPaymentsUseCase handles listing payment history.
type ListPaymentsUseCase struct {
	paymentRepo repository.PaymentRepository
}

// NewListPaymentsUseCase creates a new ListPaymentsUseCase.
func NewListPaymentsUseCase(paymentRepo repository.PaymentRepository) *ListPaymentsUseCase {
	return &ListPaymentsUseCase{paymentRepo: paymentRepo}
}

// Execute lists a user's payments with pagination.
func (uc *ListPaymentsUseCase) Execute(ctx context.Context, userID string, req *dto.ListPaymentsRequest) (*dto.ListPaymentsResponse, error) {
	p := pagination.Pagination{Page: req.Page, Limit: req.Limit}
	p.Normalize()

	payments, err := uc.paymentRepo.ListByUser(ctx, userID, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}

	count, err := uc.paymentRepo.CountByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.PaymentResponse, 0, len(payments))
	for _, pay := range payments {
		responses = append(responses, &dto.PaymentResponse{
			ID:             pay.ID,
			UserID:         pay.UserID,
			SubscriptionID: pay.SubscriptionID,
			Amount:         pay.Amount,
			Currency:       pay.Currency,
			Status:         string(pay.Status),
			Provider:       pay.Provider,
			CreatedAt:      pay.CreatedAt,
		})
	}

	return &dto.ListPaymentsResponse{
		Payments:   responses,
		TotalCount: count,
		Page:       p.Page,
		Limit:      p.Limit,
	}, nil
}
