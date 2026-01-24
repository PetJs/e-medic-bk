// Package worker provides background job workers.
package worker

import (
	"context"
	"encoding/json"
)

// EmailNotificationPayload is the payload for email notification jobs.
type EmailNotificationPayload struct {
	Type      string                 `json:"type"`
	Recipient string                 `json:"recipient"`
	Data      map[string]interface{} `json:"data"`
}

func handleSendEmail(ctx context.Context, payload json.RawMessage) error {
	var p EmailNotificationPayload
	if err := json.Unmarshal(payload, &p); err != nil {
		return err
	}

	// TODO: Send email based on type
	return nil
}
