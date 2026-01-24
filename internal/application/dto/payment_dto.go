// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// InitiatePaymentRequest represents a payment initiation request.
type InitiatePaymentRequest struct {
	Amount   int64  `json:"amount" validate:"required,min=1"`
	Currency string `json:"currency" validate:"required"`
	PlanID   string `json:"plan_id,omitempty"`
}

// PaymentIntentResponse represents a payment intent response.
type PaymentIntentResponse struct {
	ID           string `json:"id"`
	ClientSecret string `json:"client_secret"`
	Amount       int64  `json:"amount"`
	Currency     string `json:"currency"`
	Status       string `json:"status"`
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
