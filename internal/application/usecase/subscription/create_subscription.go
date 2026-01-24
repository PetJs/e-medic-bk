// Package subscription contains subscription use cases.
package subscription

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// CreateSubscriptionUseCase handles creating a subscription.
type CreateSubscriptionUseCase struct{}

func NewCreateSubscriptionUseCase() *CreateSubscriptionUseCase { return &CreateSubscriptionUseCase{} }

func (uc *CreateSubscriptionUseCase) Execute(ctx context.Context, userID string, req *dto.CreateSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	return nil, nil
}
