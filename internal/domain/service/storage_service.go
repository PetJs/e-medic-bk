// Package service defines domain service interfaces for external services.
package service

import (
	"context"
	"io"
	"time"
)

// StorageService defines the interface for object storage operations.
type StorageService interface {
	Upload(ctx context.Context, key string, body io.Reader, contentType string, size int64) error
	Download(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
	GetSignedURL(ctx context.Context, key string, expiry time.Duration) (string, error)
	GetUploadURL(ctx context.Context, key, contentType string, expiry time.Duration) (string, error)
}
