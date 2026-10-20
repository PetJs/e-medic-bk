// Package subscription contains subscription use cases.
package subscription

import (
	"context"

	"emedic-bk/internal/domain/repository"
)

// ExpireSubscriptionsUseCase reverts subscriptions to free access once the
// grace period after their paid period has elapsed.
type ExpireSubscriptionsUseCase struct {
	subRepo repository.SubscriptionRepository
}

// NewExpireSubscriptionsUseCase creates a new ExpireSubscriptionsUseCase.
func NewExpireSubscriptionsUseCase(subRepo repository.SubscriptionRepository) *ExpireSubscriptionsUseCase {
	return &ExpireSubscriptionsUseCase{subRepo: subRepo}
}

// Execute marks every stale subscription expired and returns how many rows
// changed. Access checks are grace-aware on their own; this sweep keeps the
// stored status in sync so clients can offer a renewal payment.
func (uc *ExpireSubscriptionsUseCase) Execute(ctx context.Context) (int64, error) {
	return uc.subRepo.ExpireStale(ctx)
}
