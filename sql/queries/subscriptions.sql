-- name: GetActiveSubscriptionByUser :one
SELECT * FROM subscriptions WHERE user_id = $1 AND status = 'active' LIMIT 1;

-- name: ListSubscriptionsByUser :many
SELECT * FROM subscriptions WHERE user_id = $1 ORDER BY created_at DESC;

-- name: CreateSubscription :one
INSERT INTO subscriptions (user_id, plan_id, status, current_period_start, current_period_end)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateSubscriptionStatus :one
UPDATE subscriptions SET status = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: CancelSubscription :one
UPDATE subscriptions SET status = 'canceled', canceled_at = NOW(), updated_at = NOW() WHERE id = $1 RETURNING *;
