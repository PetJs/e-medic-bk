// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// InitiatePaymentRequest represents a payment initiation request.
type InitiatePaymentRequest struct {
	PlanID string `json:"plan_id" binding:"required"`
}

// InitiateCheckoutResponse points the client at the gateway's hosted checkout.
type InitiateCheckoutResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	Reference        string `json:"reference"`
}

// VerifyPaymentRequest asks the API to confirm a transaction by reference.
type VerifyPaymentRequest struct {
	Reference string `json:"reference" binding:"required"`
}

// VerifyPaymentResponse reports the verified payment outcome.
type VerifyPaymentResponse struct {
	Status       string                `json:"status"` // "success" or "failed"
	Subscription *SubscriptionResponse `json:"subscription,omitempty"`
}

// PlanResponse describes a purchasable subscription plan.
type PlanResponse struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Amount   int64    `json:"amount"` // smallest currency unit
	Currency string   `json:"currency"`
	Interval string   `json:"interval"`
	Features []string `json:"features"`
}

// PaymentResponse represents a payment in API responses.
type PaymentResponse struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	SubscriptionID *string   `json:"subscription_id,omitempty"`
	Amount         int64     `json:"amount"`
	Currency       string    `json:"currency"`
	Status         string    `json:"status"`
	Provider       string    `json:"provider"`
	CreatedAt      time.Time `json:"created_at"`
}

// ListPaymentsRequest represents a list payments request.
type ListPaymentsRequest struct {
	Page  int `json:"page" validate:"min=1"`
	Limit int `json:"limit" validate:"min=1,max=100"`
}

// ListPaymentsResponse represents a list payments response.
type ListPaymentsResponse struct {
	Payments   []*PaymentResponse `json:"payments"`
	TotalCount int64              `json:"total_count"`
	Page       int                `json:"page"`
	Limit      int                `json:"limit"`
}

// AdminStatsResponse aggregates platform stats for the admin dashboard.
type AdminStatsResponse struct {
	TotalStudents       int64  `json:"total_students"`
	ActiveSubscriptions int64  `json:"active_subscriptions"`
	MonthlyRevenue      int64  `json:"monthly_revenue"` // smallest currency unit, last 30 days
	Currency            string `json:"currency"`
}
