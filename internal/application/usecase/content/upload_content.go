// Package content contains content management use cases.
package content

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UploadContentUseCase handles uploading PDFs and videos.
type UploadContentUseCase struct{}

func NewUploadContentUseCase() *UploadContentUseCase { return &UploadContentUseCase{} }

func (uc *UploadContentUseCase) Execute(ctx context.Context, req *dto.UploadContentRequest) (*dto.ContentResponse, error) {
	// TODO: Validate file type
	// TODO: Generate storage key
	// TODO: Upload to S3
	// TODO: Save content metadata
	return nil, nil
}
