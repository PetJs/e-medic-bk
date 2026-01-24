// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreateSubscriptionRequest represents a subscription creation request.
type CreateSubscriptionRequest struct {
	PlanID string `json:"plan_id" validate:"required"`
}

// SubscriptionResponse represents a subscription in API responses.
type SubscriptionResponse struct {
	ID                 string     `json:"id"`
	UserID             string     `json:"user_id"`
	PlanID             string     `json:"plan_id"`
	Status             string     `json:"status"`
	CurrentPeriodStart time.Time  `json:"current_period_start"`
	CurrentPeriodEnd   time.Time  `json:"current_period_end"`
	CanceledAt         *time.Time `json:"canceled_at,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
}
