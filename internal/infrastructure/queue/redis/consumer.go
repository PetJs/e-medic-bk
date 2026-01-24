// Package redis provides Redis job queue implementation.
package redis

import (
	"context"
	"encoding/json"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Consumer consumes jobs from a Redis queue.
type Consumer struct {
	client *goredis.Client
}

// NewConsumer creates a new job consumer.
func NewConsumer(client *goredis.Client) *Consumer {
	return &Consumer{client: client}
}

// JobHandler is a function that handles a job.
type JobHandler func(ctx context.Context, job *Job) error

// Consume starts consuming jobs from the queue.
func (c *Consumer) Consume(ctx context.Context, queueName string, handler JobHandler) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			result, err := c.client.BRPop(ctx, 5*time.Second, queueName).Result()
			if err != nil {
				continue // Timeout, try again
			}

			if len(result) < 2 {
				continue
			}

			var job Job
			if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
				continue // Skip invalid jobs
			}

			if err := handler(ctx, &job); err != nil {
				// TODO: Handle failed jobs (retry queue, dead letter queue)
				continue
			}
		}
	}
}
