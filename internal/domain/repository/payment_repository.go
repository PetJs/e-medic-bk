// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"
	"time"

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
	CountByUser(ctx context.Context, userID string) (int64, error)
	// SumCompletedSince returns total completed payment volume since a point in time.
	SumCompletedSince(ctx context.Context, since time.Time) (int64, error)
}
