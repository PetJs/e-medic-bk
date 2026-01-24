// Package payment contains payment use cases.
package payment

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// InitiatePaymentUseCase handles initiating a payment.
type InitiatePaymentUseCase struct{}

func NewInitiatePaymentUseCase() *InitiatePaymentUseCase { return &InitiatePaymentUseCase{} }

func (uc *InitiatePaymentUseCase) Execute(ctx context.Context, userID string, req *dto.InitiatePaymentRequest) (*dto.PaymentIntentResponse, error) {
	return nil, nil
}
