// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreateSubscriptionRequest represents a subscription creation request.
type CreateSubscriptionRequest struct {
	PlanID string `json:"plan_id" validate:"required"`
}

// SubscriptionResponse represents a subscription in API responses.
type SubscriptionResponse struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	PlanID             string    `json:"plan_id"`
	Status             string    `json:"status"`
	CurrentPeriodStart time.Time `json:"current_period_start"`
	CurrentPeriodEnd   time.Time `json:"current_period_end"`
	// GracePeriodEnd is the last moment access is granted while a renewal
	// payment is retried after CurrentPeriodEnd.
	GracePeriodEnd time.Time `json:"grace_period_end"`
	// InGrace is true once CurrentPeriodEnd has passed but access is still
	// being granted during the grace period.
	InGrace    bool       `json:"in_grace"`
	CanceledAt *time.Time `json:"canceled_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}
