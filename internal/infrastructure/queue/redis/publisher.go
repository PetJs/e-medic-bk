// Package redis provides Redis job queue implementation.
package redis

import (
	"context"
	"encoding/json"

	goredis "github.com/redis/go-redis/v9"
)

// Publisher publishes jobs to a Redis queue.
type Publisher struct {
	client *goredis.Client
}

// NewPublisher creates a new job publisher.
func NewPublisher(client *goredis.Client) *Publisher {
	return &Publisher{client: client}
}

// Job represents a background job.
type Job struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// Publish publishes a job to the queue.
func (p *Publisher) Publish(ctx context.Context, queueName string, job *Job) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return p.client.LPush(ctx, queueName, data).Err()
}
