// Package s3 provides S3-compatible object storage implementation.
package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Client wraps the S3 client.
type Client struct {
	client     *s3.Client
	bucketName string
}

// NewClient creates a new S3 client.
func NewClient(client *s3.Client, bucketName string) *Client {
	return &Client{
		client:     client,
		bucketName: bucketName,
	}
}

// Ping checks the connection to S3.
func (c *Client) Ping(ctx context.Context) error {
	_, err := c.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &c.bucketName,
	})
	return err
}
