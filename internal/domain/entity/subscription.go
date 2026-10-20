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

// SubscriptionGracePeriod is how long premium access continues after
// CurrentPeriodEnd while the renewal payment is retried. Once it elapses the
// subscription is marked expired and the user reverts to free access.
const SubscriptionGracePeriod = 72 * time.Hour

// Subscription represents a user's premium subscription.
type Subscription struct {
	ID                 string
	UserID             string
	PlanID             string
	Status             SubscriptionStatus
	CurrentPeriodStart time.Time
	CurrentPeriodEnd   time.Time
	CanceledAt         *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// AccessEndsAt is the last moment the subscription grants premium access:
// the end of the paid period plus the renewal grace period.
func (s *Subscription) AccessEndsAt() time.Time {
	return s.CurrentPeriodEnd.Add(SubscriptionGracePeriod)
}

// HasAccess reports whether the subscription grants premium access at now.
func (s *Subscription) HasAccess(now time.Time) bool {
	return s.Status == SubscriptionStatusActive && now.Before(s.AccessEndsAt())
}

// IsInGrace reports whether the paid period has ended but the grace period
// (during which renewal is retried) has not.
func (s *Subscription) IsInGrace(now time.Time) bool {
	return s.Status == SubscriptionStatusActive &&
		!now.Before(s.CurrentPeriodEnd) &&
		now.Before(s.AccessEndsAt())
}
