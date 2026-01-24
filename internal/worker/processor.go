// Package worker provides background job workers.
package worker

import (
	"context"

	queue "emedic-bk/internal/infrastructure/queue/redis"
)

// Processor processes background jobs.
type Processor struct {
	consumer *queue.Consumer
}

// NewProcessor creates a new job processor.
func NewProcessor(consumer *queue.Consumer) *Processor {
	return &Processor{consumer: consumer}
}

// Start starts processing jobs.
func (p *Processor) Start(ctx context.Context) error {
	return p.consumer.Consume(ctx, "jobs", func(ctx context.Context, job *queue.Job) error {
		switch job.Type {
		case "subscription_expiry_check":
			return handleSubscriptionExpiryCheck(ctx, job.Payload)
		case "send_email":
			return handleSendEmail(ctx, job.Payload)
		case "content_cleanup":
			return handleContentCleanup(ctx, job.Payload)
		default:
			return nil
		}
	})
}
