-- name: ListPaymentsByUser :many
SELECT * FROM payments WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: GetPaymentByProviderID :one
SELECT * FROM payments WHERE provider_payment_id = $1;

-- name: CreatePayment :one
INSERT INTO payments (user_id, subscription_id, amount, currency, status, provider, provider_payment_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdatePaymentStatus :one
UPDATE payments SET status = $2, updated_at = NOW() WHERE id = $1 RETURNING *;
