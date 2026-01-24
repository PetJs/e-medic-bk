// Package s3 provides S3-compatible object storage implementation.
package s3

import (
	"context"
	"io"
	"time"

	"emedic-bk/internal/domain/service"
)

// StorageService implements service.StorageService.
type StorageService struct {
	client *Client
}

// NewStorageService creates a new S3 storage service.
func NewStorageService(client *Client) service.StorageService {
	return &StorageService{client: client}
}

func (s *StorageService) Upload(ctx context.Context, key string, body io.Reader, contentType string, size int64) error {
	// TODO: Implement S3 upload
	return nil
}

func (s *StorageService) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	// TODO: Implement S3 download
	return nil, nil
}

func (s *StorageService) Delete(ctx context.Context, key string) error {
	// TODO: Implement S3 delete
	return nil
}

func (s *StorageService) GetSignedURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	// TODO: Implement presigned URL generation
	return "", nil
}

func (s *StorageService) GetUploadURL(ctx context.Context, key, contentType string, expiry time.Duration) (string, error) {
	// TODO: Implement presigned upload URL generation
	return "", nil
}
