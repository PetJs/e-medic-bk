// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type PaymentRepository struct{ db *DB }

func NewPaymentRepository(db *DB) repository.PaymentRepository { return &PaymentRepository{db: db} }

const paymentColumns = `id, user_id, subscription_id, amount, currency, status, provider, COALESCE(provider_payment_id, ''), created_at, updated_at`

func (r *PaymentRepository) Create(ctx context.Context, payment *entity.Payment) error {
	query := `
		INSERT INTO payments (id, user_id, subscription_id, amount, currency, status, provider, provider_payment_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		payment.ID,
		payment.UserID,
		payment.SubscriptionID,
		payment.Amount,
		payment.Currency,
		string(payment.Status),
		payment.Provider,
		payment.ProviderPaymentID,
		payment.CreatedAt,
		payment.UpdatedAt,
	)
	return err
}

func (r *PaymentRepository) GetByID(ctx context.Context, id string) (*entity.Payment, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+paymentColumns+` FROM payments WHERE id = $1`, id)
	return r.scanPayment(row)
}

func (r *PaymentRepository) GetByProviderID(ctx context.Context, providerPaymentID string) (*entity.Payment, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+paymentColumns+` FROM payments WHERE provider_payment_id = $1`, providerPaymentID)
	return r.scanPayment(row)
}

func (r *PaymentRepository) Update(ctx context.Context, payment *entity.Payment) error {
	query := `
		UPDATE payments
		SET subscription_id = $2, amount = $3, currency = $4, status = $5, provider_payment_id = $6, updated_at = $7
		WHERE id = $1
	`
	_, err := r.db.Pool.Exec(ctx, query,
		payment.ID,
		payment.SubscriptionID,
		payment.Amount,
		payment.Currency,
		string(payment.Status),
		payment.ProviderPaymentID,
		time.Now(),
	)
	return err
}

func (r *PaymentRepository) ListByUser(ctx context.Context, userID string, limit, offset int) ([]*entity.Payment, error) {
	query := `SELECT ` + paymentColumns + ` FROM payments WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Pool.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanPayments(rows)
}

func (r *PaymentRepository) ListBySubscription(ctx context.Context, subscriptionID string) ([]*entity.Payment, error) {
	query := `SELECT ` + paymentColumns + ` FROM payments WHERE subscription_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.Pool.Query(ctx, query, subscriptionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanPayments(rows)
}

func (r *PaymentRepository) CountByUser(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM payments WHERE user_id = $1`, userID).Scan(&count)
	return count, err
}

func (r *PaymentRepository) SumCompletedSince(ctx context.Context, since time.Time) (int64, error) {
	var sum int64
	err := r.db.Pool.QueryRow(ctx,
		`SELECT COALESCE(SUM(amount), 0) FROM payments WHERE status = 'completed' AND created_at >= $1`,
		since,
	).Scan(&sum)
	return sum, err
}

func (r *PaymentRepository) RevenueByDay(ctx context.Context, since time.Time) ([]entity.DailyMetric, error) {
	query := `
		SELECT DATE_TRUNC('day', created_at) AS day, COALESCE(SUM(amount), 0)
		FROM payments
		WHERE status = 'completed' AND created_at >= $1
		GROUP BY day
		ORDER BY day
	`
	rows, err := r.db.Pool.Query(ctx, query, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanDailyMetrics(rows)
}

func (r *PaymentRepository) scanPayment(row pgx.Row) (*entity.Payment, error) {
	payment := &entity.Payment{}
	var status string
	err := row.Scan(
		&payment.ID,
		&payment.UserID,
		&payment.SubscriptionID,
		&payment.Amount,
		&payment.Currency,
		&status,
		&payment.Provider,
		&payment.ProviderPaymentID,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	payment.Status = entity.PaymentStatus(status)
	return payment, nil
}

func (r *PaymentRepository) scanPayments(rows pgx.Rows) ([]*entity.Payment, error) {
	var payments []*entity.Payment
	for rows.Next() {
		payment := &entity.Payment{}
		var status string
		err := rows.Scan(
			&payment.ID,
			&payment.UserID,
			&payment.SubscriptionID,
			&payment.Amount,
			&payment.Currency,
			&status,
			&payment.Provider,
			&payment.ProviderPaymentID,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payment.Status = entity.PaymentStatus(status)
		payments = append(payments, payment)
	}
	return payments, rows.Err()
}
