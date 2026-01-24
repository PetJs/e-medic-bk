// Package entity contains the core domain entities.
package entity

import "time"

// SubscriptionStatus represents the status of a subscription.
type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusCanceled SubscriptionStatus = "canceled"
	SubscriptionStatusExpired  SubscriptionStatus = "expired"
	SubscriptionStatusPending  SubscriptionStatus = "pending"
)

// Subscription represents a user's premium subscription.
type Subscription struct {
	ID              string
	UserID          string
	PlanID          string
	Status          SubscriptionStatus
	CurrentPeriodStart time.Time
	CurrentPeriodEnd   time.Time
	CanceledAt      *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
