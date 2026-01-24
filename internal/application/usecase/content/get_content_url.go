// Package content contains content management use cases.
package content

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// GetContentURLUseCase handles generating signed URLs for content access.
type GetContentURLUseCase struct{}

func NewGetContentURLUseCase() *GetContentURLUseCase { return &GetContentURLUseCase{} }

func (uc *GetContentURLUseCase) Execute(ctx context.Context, contentID, userID string) (*dto.ContentURLResponse, error) {
	// TODO: Get content metadata
	// TODO: Check user access (enrollment, subscription)
	// TODO: Generate signed URL
	return nil, nil
}
