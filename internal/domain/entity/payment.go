// Package entity contains the core domain entities.
package entity

import "time"

// PaymentStatus represents the status of a payment.
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

// Payment represents a payment transaction.
type Payment struct {
	ID              string
	UserID          string
	SubscriptionID  *string
	Amount          int64  // in smallest currency unit (e.g., cents)
	Currency        string // ISO 4217 code
	Status          PaymentStatus
	Provider        string // "stripe", "paystack"
	ProviderPaymentID string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
