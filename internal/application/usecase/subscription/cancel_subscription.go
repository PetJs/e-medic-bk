// Package subscription contains subscription use cases.
package subscription

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/domain/entity"
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

// Execute cancels a user's subscription.
func (uc *CancelSubscriptionUseCase) Execute(ctx context.Context, userID, subscriptionID string) error {
	sub, err := uc.subRepo.GetByID(ctx, subscriptionID)
	if err != nil {
		return err
	}
	if sub == nil || sub.UserID != userID {
		return ErrSubscriptionNotFound
	}

	now := time.Now()
	sub.Status = entity.SubscriptionStatusCanceled
	sub.CanceledAt = &now
	sub.UpdatedAt = now
	return uc.subRepo.Update(ctx, sub)
}
