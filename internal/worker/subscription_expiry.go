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
	// Expiry is enforced by the API server's periodic sweep
	// (subscription.ExpireSubscriptionsUseCase) and by grace-aware access
	// queries. This worker job is kept for deployments that later move the
	// sweep off the API process.
	return nil
}
