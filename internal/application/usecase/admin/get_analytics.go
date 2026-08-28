package admin

import (
	"context"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// DefaultAnalyticsDays is used when the caller doesn't specify a range.
const DefaultAnalyticsDays = 30

// GetAnalyticsUseCase aggregates trend data for the admin analytics dashboard.
type GetAnalyticsUseCase struct {
	paymentRepo  repository.PaymentRepository
	userRepo     repository.UserRepository
	subRepo      repository.SubscriptionRepository
	progressRepo repository.ProgressRepository
	currency     string
}

// NewGetAnalyticsUseCase creates a new GetAnalyticsUseCase.
func NewGetAnalyticsUseCase(
	paymentRepo repository.PaymentRepository,
	userRepo repository.UserRepository,
	subRepo repository.SubscriptionRepository,
	progressRepo repository.ProgressRepository,
	currency string,
) *GetAnalyticsUseCase {
	return &GetAnalyticsUseCase{
		paymentRepo:  paymentRepo,
		userRepo:     userRepo,
		subRepo:      subRepo,
		progressRepo: progressRepo,
		currency:     currency,
	}
}

// Execute returns revenue/signup/new-subscription trends over the last `days`
// days, plus a current-state per-module completion-rate breakdown (not
// date-scoped — it's a snapshot, not a trend).
func (uc *GetAnalyticsUseCase) Execute(ctx context.Context, days int) (*dto.AdminAnalyticsResponse, error) {
	if days <= 0 {
		days = DefaultAnalyticsDays
	}
	since := time.Now().AddDate(0, 0, -days)

	revenue, err := uc.paymentRepo.RevenueByDay(ctx, since)
	if err != nil {
		return nil, err
	}
	signups, err := uc.userRepo.SignupsByDay(ctx, since)
	if err != nil {
		return nil, err
	}
	newSubs, err := uc.subRepo.NewSubscriptionsByDay(ctx, since)
	if err != nil {
		return nil, err
	}
	completionRates, err := uc.progressRepo.GetModuleCompletionRates(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.AdminAnalyticsResponse{
		RevenueTrend:          toDailyMetricResponses(revenue),
		SignupTrend:           toDailyMetricResponses(signups),
		SubscriptionTrend:     toDailyMetricResponses(newSubs),
		ModuleCompletionRates: toModuleCompletionResponses(completionRates),
		Currency:              uc.currency,
	}, nil
}

func toDailyMetricResponses(metrics []entity.DailyMetric) []dto.DailyMetricResponse {
	responses := make([]dto.DailyMetricResponse, 0, len(metrics))
	for _, m := range metrics {
		responses = append(responses, dto.DailyMetricResponse{
			Date:  m.Date.Format("2006-01-02"),
			Value: m.Value,
		})
	}
	return responses
}

func toModuleCompletionResponses(rates []entity.ModuleCompletionRate) []dto.ModuleCompletionRateResponse {
	responses := make([]dto.ModuleCompletionRateResponse, 0, len(rates))
	for _, r := range rates {
		responses = append(responses, dto.ModuleCompletionRateResponse{
			ModuleID:      r.ModuleID,
			ModuleTitle:   r.ModuleTitle,
			CompletionPct: r.CompletionPct,
		})
	}
	return responses
}
