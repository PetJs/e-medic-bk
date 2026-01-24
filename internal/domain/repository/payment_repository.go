// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// PaymentRepository defines the interface for payment data access.
type PaymentRepository interface {
	Create(ctx context.Context, payment *entity.Payment) error
	GetByID(ctx context.Context, id string) (*entity.Payment, error)
	GetByProviderID(ctx context.Context, providerPaymentID string) (*entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*entity.Payment, error)
	ListBySubscription(ctx context.Context, subscriptionID string) ([]*entity.Payment, error)
}
