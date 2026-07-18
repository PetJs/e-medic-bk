// Package payment contains payment use cases.
package payment

import (
	"context"
	"errors"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/subscription"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// ErrPaymentNotFound is returned when no payment matches the reference.
var ErrPaymentNotFound = errors.New("payment not found")

// VerifyPaymentUseCase confirms a transaction with the gateway and, on
// success, marks the payment completed and activates the subscription.
// It is idempotent: re-verifying a completed payment is a no-op success.
type VerifyPaymentUseCase struct {
	paymentRepo repository.PaymentRepository
	gateway     service.PaymentService
	activateUC  *subscription.ActivateSubscriptionUseCase
}

// NewVerifyPaymentUseCase creates a new VerifyPaymentUseCase.
func NewVerifyPaymentUseCase(
	paymentRepo repository.PaymentRepository,
	gateway service.PaymentService,
	activateUC *subscription.ActivateSubscriptionUseCase,
) *VerifyPaymentUseCase {
	return &VerifyPaymentUseCase{
		paymentRepo: paymentRepo,
		gateway:     gateway,
		activateUC:  activateUC,
	}
}

// Execute verifies the transaction behind a reference.
func (uc *VerifyPaymentUseCase) Execute(ctx context.Context, reference string) (*dto.VerifyPaymentResponse, error) {
	payment, err := uc.paymentRepo.GetByProviderID(ctx, reference)
	if err != nil {
		return nil, err
	}
	if payment == nil {
		return nil, ErrPaymentNotFound
	}

	// Already processed — report success without re-activating.
	if payment.Status == entity.PaymentStatusCompleted {
		return &dto.VerifyPaymentResponse{Status: "success"}, nil
	}

	status, err := uc.gateway.VerifyTransaction(ctx, reference)
	if err != nil {
		return nil, err
	}

	switch status.Status {
	case "success":
		// proceed to activation below
	case "failed", "reversed":
		payment.Status = entity.PaymentStatusFailed
		if err := uc.paymentRepo.Update(ctx, payment); err != nil {
			return nil, err
		}
		return &dto.VerifyPaymentResponse{Status: "failed"}, nil
	default:
		// abandoned / pending / ongoing / processing / queued — the transaction
		// can still complete; leave the payment untouched so a later verify or
		// the webhook can settle it.
		return &dto.VerifyPaymentResponse{Status: "pending"}, nil
	}

	sub, err := uc.activateUC.Execute(ctx, payment.UserID, PlanMonthly)
	if err != nil {
		return nil, err
	}

	payment.Status = entity.PaymentStatusCompleted
	payment.SubscriptionID = &sub.ID
	if err := uc.paymentRepo.Update(ctx, payment); err != nil {
		return nil, err
	}

	return &dto.VerifyPaymentResponse{Status: "success", Subscription: sub}, nil
}
