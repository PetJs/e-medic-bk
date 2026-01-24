// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// SubscriptionRepository defines the interface for subscription data access.
type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *entity.Subscription) error
	GetByID(ctx context.Context, id string) (*entity.Subscription, error)
	GetActiveByUser(ctx context.Context, userID string) (*entity.Subscription, error)
	Update(ctx context.Context, subscription *entity.Subscription) error
	Delete(ctx context.Context, id string) error
	ListByUser(ctx context.Context, userID string) ([]*entity.Subscription, error)
	ListExpiring(ctx context.Context, withinDays int) ([]*entity.Subscription, error)
}
