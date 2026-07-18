// Package admin contains admin-facing use cases.
package admin

import (
	"context"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// GetStatsUseCase aggregates platform stats for the admin dashboard.
type GetStatsUseCase struct {
	userRepo    repository.UserRepository
	subRepo     repository.SubscriptionRepository
	paymentRepo repository.PaymentRepository
	currency    string
}

// NewGetStatsUseCase creates a new GetStatsUseCase.
func NewGetStatsUseCase(
	userRepo repository.UserRepository,
	subRepo repository.SubscriptionRepository,
	paymentRepo repository.PaymentRepository,
	currency string,
) *GetStatsUseCase {
	return &GetStatsUseCase{
		userRepo:    userRepo,
		subRepo:     subRepo,
		paymentRepo: paymentRepo,
		currency:    currency,
	}
}

// Execute returns totals for students, active subscriptions, and 30-day revenue.
func (uc *GetStatsUseCase) Execute(ctx context.Context) (*dto.AdminStatsResponse, error) {
	students, err := uc.userRepo.CountByRole(ctx, "student")
	if err != nil {
		return nil, err
	}

	activeSubs, err := uc.subRepo.CountActive(ctx)
	if err != nil {
		return nil, err
	}

	revenue, err := uc.paymentRepo.SumCompletedSince(ctx, time.Now().AddDate(0, 0, -30))
	if err != nil {
		return nil, err
	}

	return &dto.AdminStatsResponse{
		TotalStudents:       students,
		ActiveSubscriptions: activeSubs,
		MonthlyRevenue:      revenue,
		Currency:            uc.currency,
	}, nil
}
