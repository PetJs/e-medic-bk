// Package s3 provides S3-compatible object storage implementation.
package s3

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"

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
	_, err := s.client.client.PutObject(ctx, &awss3.PutObjectInput{
		Bucket:        &s.client.bucketName,
		Key:           &key,
		Body:          body,
		ContentType:   &contentType,
		ContentLength: aws.Int64(size),
	})
	return err
}

func (s *StorageService) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	out, err := s.client.client.GetObject(ctx, &awss3.GetObjectInput{
		Bucket: &s.client.bucketName,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	return out.Body, nil
}

func (s *StorageService) Delete(ctx context.Context, key string) error {
	_, err := s.client.client.DeleteObject(ctx, &awss3.DeleteObjectInput{
		Bucket: &s.client.bucketName,
		Key:    &key,
	})
	return err
}

func (s *StorageService) GetSignedURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	req, err := s.client.presign.PresignGetObject(ctx, &awss3.GetObjectInput{
		Bucket: &s.client.bucketName,
		Key:    &key,
	}, awss3.WithPresignExpires(expiry))
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *StorageService) GetUploadURL(ctx context.Context, key, contentType string, expiry time.Duration) (string, error) {
	req, err := s.client.presign.PresignPutObject(ctx, &awss3.PutObjectInput{
		Bucket:      &s.client.bucketName,
		Key:         &key,
		ContentType: &contentType,
	}, awss3.WithPresignExpires(expiry))
	if err != nil {
		return "", err
	}
	return req.URL, nil
}
