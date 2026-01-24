// Package subscription contains subscription use cases.
package subscription

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// RenewSubscriptionUseCase handles renewing a subscription.
type RenewSubscriptionUseCase struct{}

func NewRenewSubscriptionUseCase() *RenewSubscriptionUseCase { return &RenewSubscriptionUseCase{} }

func (uc *RenewSubscriptionUseCase) Execute(ctx context.Context, userID, subscriptionID string) (*dto.SubscriptionResponse, error) {
	return nil, nil
}
