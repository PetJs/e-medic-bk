// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type SubscriptionRepository struct{ db *DB }

func NewSubscriptionRepository(db *DB) repository.SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(ctx context.Context, subscription *entity.Subscription) error {
	return nil
}
func (r *SubscriptionRepository) GetByID(ctx context.Context, id string) (*entity.Subscription, error) {
	return nil, nil
}
func (r *SubscriptionRepository) GetActiveByUser(ctx context.Context, userID string) (*entity.Subscription, error) {
	return nil, nil
}
func (r *SubscriptionRepository) Update(ctx context.Context, subscription *entity.Subscription) error {
	return nil
}
func (r *SubscriptionRepository) Delete(ctx context.Context, id string) error { return nil }
func (r *SubscriptionRepository) ListByUser(ctx context.Context, userID string) ([]*entity.Subscription, error) {
	return nil, nil
}
func (r *SubscriptionRepository) ListExpiring(ctx context.Context, withinDays int) ([]*entity.Subscription, error) {
	return nil, nil
}
