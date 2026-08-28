// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/domain/entity"
)

// ErrDuplicateActiveSubscription is returned by Create when the user already
// has an active subscription (a concurrent activation won the race).
var ErrDuplicateActiveSubscription = errors.New("user already has an active subscription")

// SubscriptionRepository defines the interface for subscription data access.
type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *entity.Subscription) error
	GetByID(ctx context.Context, id string) (*entity.Subscription, error)
	GetActiveByUser(ctx context.Context, userID string) (*entity.Subscription, error)
	Update(ctx context.Context, subscription *entity.Subscription) error
	Delete(ctx context.Context, id string) error
	ListByUser(ctx context.Context, userID string) ([]*entity.Subscription, error)
	ListExpiring(ctx context.Context, withinDays int) ([]*entity.Subscription, error)
	CountActive(ctx context.Context) (int64, error)
	// NewSubscriptionsByDay returns new subscriptions created grouped by day, since a point in time.
	NewSubscriptionsByDay(ctx context.Context, since time.Time) ([]entity.DailyMetric, error)
}
