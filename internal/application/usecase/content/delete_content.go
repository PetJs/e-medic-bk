// Package content contains content management use cases.
package content

import (
	"context"

	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// DeleteContentUseCase handles deleting content.
type DeleteContentUseCase struct {
	contentRepo repository.ContentRepository
	storage     service.StorageService
}

// NewDeleteContentUseCase creates a new DeleteContentUseCase.
func NewDeleteContentUseCase(contentRepo repository.ContentRepository, storage service.StorageService) *DeleteContentUseCase {
	return &DeleteContentUseCase{contentRepo: contentRepo, storage: storage}
}

// Execute removes content metadata and its stored object.
func (uc *DeleteContentUseCase) Execute(ctx context.Context, contentID string) error {
	content, err := uc.contentRepo.GetByID(ctx, contentID)
	if err != nil {
		return err
	}
	if content == nil {
		return ErrContentNotFound
	}

	if err := uc.contentRepo.Delete(ctx, contentID); err != nil {
		return err
	}
	// Object cleanup is best effort; the DB row is the source of truth.
	_ = uc.storage.Delete(ctx, content.StorageKey)
	return nil
}
