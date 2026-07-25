// Package service defines domain service interfaces for external services.
package service

import "context"

// ImageGenerator generates an image from a text prompt.
type ImageGenerator interface {
	GenerateImage(ctx context.Context, prompt string) (data []byte, mimeType string, err error)
}
