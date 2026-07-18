// Package s3 provides S3-compatible object storage implementation.
package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Config holds S3-compatible storage settings (AWS S3, MinIO, Cloudflare R2).
type Config struct {
	Endpoint        string // empty for AWS S3
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	UsePathStyle    bool // true for MinIO
}

// Client wraps the S3 client and its presigner.
type Client struct {
	client     *s3.Client
	presign    *s3.PresignClient
	bucketName string
}

// NewClient creates a new S3 client from explicit configuration.
func NewClient(cfg Config) *Client {
	options := s3.Options{
		Region:       cfg.Region,
		UsePathStyle: cfg.UsePathStyle,
		Credentials: aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     cfg.AccessKeyID,
				SecretAccessKey: cfg.SecretAccessKey,
			}, nil
		}),
	}
	if cfg.Endpoint != "" {
		options.BaseEndpoint = aws.String(cfg.Endpoint)
	}

	client := s3.New(options)
	return &Client{
		client:     client,
		presign:    s3.NewPresignClient(client),
		bucketName: cfg.BucketName,
	}
}

// Ping checks the connection to S3.
func (c *Client) Ping(ctx context.Context) error {
	_, err := c.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &c.bucketName,
	})
	return err
}
