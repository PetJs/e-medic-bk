// Package payment contains payment use cases.
package payment

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// ErrUnknownPlan is returned when an unrecognized plan is requested.
var ErrUnknownPlan = errors.New("unknown plan")

// PlanMonthly is the single MVP plan.
const PlanMonthly = "monthly"

// PlanDetails describes the configured plan pricing.
type PlanDetails struct {
	Amount   int64
	Currency string
}

// InitiatePaymentUseCase starts a hosted checkout for a subscription plan.
type InitiatePaymentUseCase struct {
	paymentRepo repository.PaymentRepository
	userRepo    repository.UserRepository
	gateway     service.PaymentService
	idGen       port.IDGenerator
	plan        PlanDetails
	callbackURL string
}

// NewInitiatePaymentUseCase creates a new InitiatePaymentUseCase.
func NewInitiatePaymentUseCase(
	paymentRepo repository.PaymentRepository,
	userRepo repository.UserRepository,
	gateway service.PaymentService,
	idGen port.IDGenerator,
	plan PlanDetails,
	callbackURL string,
) *InitiatePaymentUseCase {
	return &InitiatePaymentUseCase{
		paymentRepo: paymentRepo,
		userRepo:    userRepo,
		gateway:     gateway,
		idGen:       idGen,
		plan:        plan,
		callbackURL: callbackURL,
	}
}

// Execute creates a pending payment and returns the gateway checkout URL.
func (uc *InitiatePaymentUseCase) Execute(ctx context.Context, userID string, req *dto.InitiatePaymentRequest) (*dto.InitiateCheckoutResponse, error) {
	if req.PlanID != PlanMonthly {
		return nil, ErrUnknownPlan
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	now := time.Now()
	payment := &entity.Payment{
		ID:        uc.idGen.Generate(),
		UserID:    userID,
		Amount:    uc.plan.Amount,
		Currency:  uc.plan.Currency,
		Status:    entity.PaymentStatusPending,
		Provider:  "paystack",
		CreatedAt: now,
		UpdatedAt: now,
	}
	// The payment ID doubles as the gateway transaction reference.
	payment.ProviderPaymentID = payment.ID

	if err := uc.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}

	session, err := uc.gateway.InitializeTransaction(
		ctx,
		user.Email,
		payment.Amount,
		payment.Currency,
		payment.ProviderPaymentID,
		uc.callbackURL,
	)
	if err != nil {
		return nil, err
	}

	return &dto.InitiateCheckoutResponse{
		AuthorizationURL: session.AuthorizationURL,
		Reference:        session.Reference,
	}, nil
}
