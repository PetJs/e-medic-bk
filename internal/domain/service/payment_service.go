// Package service defines domain service interfaces for external services.
package service

import "context"

// PaymentIntent represents an intent to collect payment.
type PaymentIntent struct {
	ID           string
	ClientSecret string
	Amount       int64
	Currency     string
	Status       string
}

// PaymentService defines the interface for payment gateway operations.
type PaymentService interface {
	CreatePaymentIntent(ctx context.Context, amount int64, currency, customerID string) (*PaymentIntent, error)
	ConfirmPayment(ctx context.Context, paymentIntentID string) (*PaymentIntent, error)
	CreateCustomer(ctx context.Context, email, name string) (customerID string, err error)
	CreateSubscription(ctx context.Context, customerID, priceID string) (subscriptionID string, err error)
	CancelSubscription(ctx context.Context, subscriptionID string) error
	HandleWebhook(ctx context.Context, payload []byte, signature string) (eventType string, data map[string]interface{}, err error)
}
