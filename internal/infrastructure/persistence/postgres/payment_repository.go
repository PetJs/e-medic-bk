// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type PaymentRepository struct{ db *DB }

func NewPaymentRepository(db *DB) repository.PaymentRepository { return &PaymentRepository{db: db} }

func (r *PaymentRepository) Create(ctx context.Context, payment *entity.Payment) error { return nil }
func (r *PaymentRepository) GetByID(ctx context.Context, id string) (*entity.Payment, error) {
	return nil, nil
}
func (r *PaymentRepository) GetByProviderID(ctx context.Context, providerPaymentID string) (*entity.Payment, error) {
	return nil, nil
}
func (r *PaymentRepository) Update(ctx context.Context, payment *entity.Payment) error { return nil }
func (r *PaymentRepository) ListByUser(ctx context.Context, userID string, limit, offset int) ([]*entity.Payment, error) {
	return nil, nil
}
func (r *PaymentRepository) ListBySubscription(ctx context.Context, subscriptionID string) ([]*entity.Payment, error) {
	return nil, nil
}
