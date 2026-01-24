// Package port defines secondary port interfaces for the application layer.
package port

import "context"

// Mailer defines the interface for sending emails.
type Mailer interface {
	SendPasswordReset(ctx context.Context, email, token string) error
	SendWelcome(ctx context.Context, email, name string) error
	SendSubscriptionConfirmation(ctx context.Context, email, planName string) error
}
