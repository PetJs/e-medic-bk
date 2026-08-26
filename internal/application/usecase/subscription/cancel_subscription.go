// Package subscription contains subscription use cases.
package subscription

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/domain/repository"
)

// ErrSubscriptionNotFound is returned when the subscription does not exist
// or belongs to another user.
var ErrSubscriptionNotFound = errors.New("subscription not found")

// CancelSubscriptionUseCase handles canceling a subscription.
type CancelSubscriptionUseCase struct {
	subRepo repository.SubscriptionRepository
}

// NewCancelSubscriptionUseCase creates a new CancelSubscriptionUseCase.
func NewCancelSubscriptionUseCase(subRepo repository.SubscriptionRepository) *CancelSubscriptionUseCase {
	return &CancelSubscriptionUseCase{subRepo: subRepo}
}

// Execute marks a subscription to not renew. Status is deliberately left
// "active" — access is driven by CurrentPeriodEnd (see
// SubscriptionRepository.GetActiveByUser), so the student keeps the access
// they already paid for and it simply won't be extended again. CanceledAt
// is the "won't renew" signal for the UI. Idempotent: canceling an
// already-canceled subscription just re-confirms it.
func (uc *CancelSubscriptionUseCase) Execute(ctx context.Context, userID, subscriptionID string) error {
	sub, err := uc.subRepo.GetByID(ctx, subscriptionID)
	if err != nil {
		return err
	}
	if sub == nil || sub.UserID != userID {
		return ErrSubscriptionNotFound
	}

	now := time.Now()
	sub.CanceledAt = &now
	sub.UpdatedAt = now
	return uc.subRepo.Update(ctx, sub)
}
