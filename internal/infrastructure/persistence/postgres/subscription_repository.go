// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type SubscriptionRepository struct{ db *DB }

func NewSubscriptionRepository(db *DB) repository.SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

const subscriptionColumns = `id, user_id, plan_id, status, current_period_start, current_period_end, canceled_at, created_at, updated_at`

func (r *SubscriptionRepository) Create(ctx context.Context, subscription *entity.Subscription) error {
	query := `
		INSERT INTO subscriptions (id, user_id, plan_id, status, current_period_start, current_period_end, canceled_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		subscription.ID,
		subscription.UserID,
		subscription.PlanID,
		string(subscription.Status),
		subscription.CurrentPeriodStart,
		subscription.CurrentPeriodEnd,
		subscription.CanceledAt,
		subscription.CreatedAt,
		subscription.UpdatedAt,
	)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "idx_subscriptions_one_active" {
		return repository.ErrDuplicateActiveSubscription
	}
	return err
}

func (r *SubscriptionRepository) GetByID(ctx context.Context, id string) (*entity.Subscription, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+subscriptionColumns+` FROM subscriptions WHERE id = $1`, id)
	return r.scanSubscription(row)
}

func (r *SubscriptionRepository) GetActiveByUser(ctx context.Context, userID string) (*entity.Subscription, error) {
	query := `
		SELECT ` + subscriptionColumns + `
		FROM subscriptions
		WHERE user_id = $1 AND status = 'active' AND current_period_end > NOW()
		ORDER BY current_period_end DESC
		LIMIT 1
	`
	row := r.db.Pool.QueryRow(ctx, query, userID)
	return r.scanSubscription(row)
}

func (r *SubscriptionRepository) Update(ctx context.Context, subscription *entity.Subscription) error {
	query := `
		UPDATE subscriptions
		SET plan_id = $2, status = $3, current_period_start = $4, current_period_end = $5, canceled_at = $6, updated_at = $7
		WHERE id = $1
	`
	_, err := r.db.Pool.Exec(ctx, query,
		subscription.ID,
		subscription.PlanID,
		string(subscription.Status),
		subscription.CurrentPeriodStart,
		subscription.CurrentPeriodEnd,
		subscription.CanceledAt,
		time.Now(),
	)
	return err
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM subscriptions WHERE id = $1`, id)
	return err
}

func (r *SubscriptionRepository) ListByUser(ctx context.Context, userID string) ([]*entity.Subscription, error) {
	query := `SELECT ` + subscriptionColumns + ` FROM subscriptions WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanSubscriptions(rows)
}

func (r *SubscriptionRepository) ListExpiring(ctx context.Context, withinDays int) ([]*entity.Subscription, error) {
	query := `
		SELECT ` + subscriptionColumns + `
		FROM subscriptions
		WHERE status = 'active' AND current_period_end BETWEEN NOW() AND NOW() + ($1 * INTERVAL '1 day')
		ORDER BY current_period_end
	`
	rows, err := r.db.Pool.Query(ctx, query, withinDays)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanSubscriptions(rows)
}

func (r *SubscriptionRepository) CountActive(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.Pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM subscriptions WHERE status = 'active' AND current_period_end > NOW()`,
	).Scan(&count)
	return count, err
}

func (r *SubscriptionRepository) NewSubscriptionsByDay(ctx context.Context, since time.Time) ([]entity.DailyMetric, error) {
	query := `
		SELECT DATE_TRUNC('day', created_at) AS day, COUNT(*)
		FROM subscriptions
		WHERE created_at >= $1
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

func (r *SubscriptionRepository) scanSubscription(row pgx.Row) (*entity.Subscription, error) {
	sub := &entity.Subscription{}
	var status string
	err := row.Scan(
		&sub.ID,
		&sub.UserID,
		&sub.PlanID,
		&status,
		&sub.CurrentPeriodStart,
		&sub.CurrentPeriodEnd,
		&sub.CanceledAt,
		&sub.CreatedAt,
		&sub.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	sub.Status = entity.SubscriptionStatus(status)
	return sub, nil
}

func (r *SubscriptionRepository) scanSubscriptions(rows pgx.Rows) ([]*entity.Subscription, error) {
	var subs []*entity.Subscription
	for rows.Next() {
		sub := &entity.Subscription{}
		var status string
		err := rows.Scan(
			&sub.ID,
			&sub.UserID,
			&sub.PlanID,
			&status,
			&sub.CurrentPeriodStart,
			&sub.CurrentPeriodEnd,
			&sub.CanceledAt,
			&sub.CreatedAt,
			&sub.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sub.Status = entity.SubscriptionStatus(status)
		subs = append(subs, sub)
	}
	return subs, rows.Err()
}
