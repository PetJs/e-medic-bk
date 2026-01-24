// Package worker provides background job workers.
package worker

import (
	"context"
	"encoding/json"
)

// ContentCleanupPayload is the payload for content cleanup jobs.
type ContentCleanupPayload struct {
	// No payload needed for periodic cleanup
}

func handleContentCleanup(ctx context.Context, payload json.RawMessage) error {
	// TODO: Find orphaned content in S3
	// TODO: Delete orphaned content
	return nil
}
