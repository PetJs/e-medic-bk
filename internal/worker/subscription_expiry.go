// Package worker provides background job workers.
package worker

import (
	"context"
	"encoding/json"
)

// SubscriptionExpiryPayload is the payload for subscription expiry jobs.
type SubscriptionExpiryPayload struct {
	// No payload needed for periodic check
}

func handleSubscriptionExpiryCheck(ctx context.Context, payload json.RawMessage) error {
	// TODO: Find subscriptions expiring soon
	// TODO: Send reminder emails
	// TODO: Mark expired subscriptions
	return nil
}
