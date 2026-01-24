// Package subscription contains subscription use cases.
package subscription

import "context"

// CancelSubscriptionUseCase handles canceling a subscription.
type CancelSubscriptionUseCase struct{}

func NewCancelSubscriptionUseCase() *CancelSubscriptionUseCase { return &CancelSubscriptionUseCase{} }

func (uc *CancelSubscriptionUseCase) Execute(ctx context.Context, userID, subscriptionID string) error {
	return nil
}
