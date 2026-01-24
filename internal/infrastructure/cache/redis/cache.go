// Package redis provides Redis cache implementation.
package redis

import (
	"context"
	"time"
)

// Cache provides caching operations using Redis.
type Cache struct {
	client *Client
}

// NewCache creates a new Redis cache.
func NewCache(client *Client) *Cache {
	return &Cache{client: client}
}

// Get retrieves a value from the cache.
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Redis().Get(ctx, key).Result()
}

// Set stores a value in the cache with an expiration.
func (c *Cache) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return c.client.Redis().Set(ctx, key, value, expiration).Err()
}

// Delete removes a value from the cache.
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Redis().Del(ctx, key).Err()
}
