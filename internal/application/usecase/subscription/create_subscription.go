// Package subscription contains subscription use cases.
package subscription

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

func toSubscriptionResponse(s *entity.Subscription) *dto.SubscriptionResponse {
	return &dto.SubscriptionResponse{
		ID:                 s.ID,
		UserID:             s.UserID,
		PlanID:             s.PlanID,
		Status:             string(s.Status),
		CurrentPeriodStart: s.CurrentPeriodStart,
		CurrentPeriodEnd:   s.CurrentPeriodEnd,
		CanceledAt:         s.CanceledAt,
		CreatedAt:          s.CreatedAt,
	}
}

// ActivateSubscriptionUseCase activates (or extends) a user's subscription
// after a confirmed payment.
type ActivateSubscriptionUseCase struct {
	subRepo repository.SubscriptionRepository
	idGen   port.IDGenerator
}

// NewActivateSubscriptionUseCase creates a new ActivateSubscriptionUseCase.
func NewActivateSubscriptionUseCase(subRepo repository.SubscriptionRepository, idGen port.IDGenerator) *ActivateSubscriptionUseCase {
	return &ActivateSubscriptionUseCase{subRepo: subRepo, idGen: idGen}
}

// Execute grants one billing period (30 days). An existing active
// subscription is extended from its current period end.
func (uc *ActivateSubscriptionUseCase) Execute(ctx context.Context, userID, planID string) (*dto.SubscriptionResponse, error) {
	const period = 30 * 24 * time.Hour
	now := time.Now()

	existing, err := uc.subRepo.GetActiveByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		existing.CurrentPeriodEnd = existing.CurrentPeriodEnd.Add(period)
		existing.UpdatedAt = now
		if err := uc.subRepo.Update(ctx, existing); err != nil {
			return nil, err
		}
		return toSubscriptionResponse(existing), nil
	}

	sub := &entity.Subscription{
		ID:                 uc.idGen.Generate(),
		UserID:             userID,
		PlanID:             planID,
		Status:             entity.SubscriptionStatusActive,
		CurrentPeriodStart: now,
		CurrentPeriodEnd:   now.Add(period),
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	if err := uc.subRepo.Create(ctx, sub); err != nil {
		if errors.Is(err, repository.ErrDuplicateActiveSubscription) {
			// A concurrent activation won the race (typically a double verify
			// of the same payment) — return the winner rather than double-crediting.
			winner, gErr := uc.subRepo.GetActiveByUser(ctx, userID)
			if gErr == nil && winner != nil {
				return toSubscriptionResponse(winner), nil
			}
		}
		return nil, err
	}
	return toSubscriptionResponse(sub), nil
}
