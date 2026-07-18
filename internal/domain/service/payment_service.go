// Package service defines domain service interfaces for external services.
package service

import (
	"context"
	"time"
)

// CheckoutSession represents a hosted checkout session (redirect flow).
type CheckoutSession struct {
	AuthorizationURL string
	AccessCode       string
	Reference        string
}

// TransactionStatus represents the verified state of a transaction.
type TransactionStatus struct {
	Reference string
	Status    string // "success", "failed", "abandoned", ...
	Amount    int64  // smallest currency unit
	Currency  string
	PaidAt    time.Time
}

// PaymentService defines the interface for payment gateway operations.
type PaymentService interface {
	// InitializeTransaction starts a hosted checkout and returns the redirect URL.
	InitializeTransaction(ctx context.Context, email string, amount int64, currency, reference, callbackURL string) (*CheckoutSession, error)
	// VerifyTransaction confirms a transaction's final status with the gateway.
	VerifyTransaction(ctx context.Context, reference string) (*TransactionStatus, error)
	// VerifyWebhookSignature checks a webhook payload's authenticity.
	VerifyWebhookSignature(payload []byte, signature string) bool
}
