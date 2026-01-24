// Package valueobject contains immutable value objects with validation.
package valueobject

import "errors"

// ErrInvalidSubscriptionStatus is returned when a subscription status is invalid.
var ErrInvalidSubscriptionStatus = errors.New("invalid subscription status")

// SubscriptionStatus represents the status of a subscription.
type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusCanceled SubscriptionStatus = "canceled"
	SubscriptionStatusExpired  SubscriptionStatus = "expired"
	SubscriptionStatusPending  SubscriptionStatus = "pending"
)

// NewSubscriptionStatus creates a new SubscriptionStatus value object with validation.
func NewSubscriptionStatus(status string) (SubscriptionStatus, error) {
	switch SubscriptionStatus(status) {
	case SubscriptionStatusActive, SubscriptionStatusCanceled, SubscriptionStatusExpired, SubscriptionStatusPending:
		return SubscriptionStatus(status), nil
	default:
		return "", ErrInvalidSubscriptionStatus
	}
}

// String returns the subscription status as a string.
func (s SubscriptionStatus) String() string {
	return string(s)
}

// IsActive returns true if the subscription is active.
func (s SubscriptionStatus) IsActive() bool {
	return s == SubscriptionStatusActive
}
