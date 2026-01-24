// Package content contains content management use cases.
package content

import "context"

// DeleteContentUseCase handles deleting content.
type DeleteContentUseCase struct{}

func NewDeleteContentUseCase() *DeleteContentUseCase { return &DeleteContentUseCase{} }

func (uc *DeleteContentUseCase) Execute(ctx context.Context, contentID string) error {
	// TODO: Get content metadata
	// TODO: Delete from S3
	// TODO: Delete metadata
	return nil
}
