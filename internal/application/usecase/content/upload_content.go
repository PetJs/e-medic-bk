// Package content contains content management use cases.
package content

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// ErrLessonNotFound is returned when the target lesson does not exist.
var ErrLessonNotFound = errors.New("lesson not found")

// ErrUnsupportedType is returned for file types other than video/PDF/image.
var ErrUnsupportedType = errors.New("unsupported content type")

// TypeFromMime maps a MIME type to a domain content type.
func TypeFromMime(mimeType string) (entity.ContentType, error) {
	switch {
	case strings.HasPrefix(mimeType, "video/"):
		return entity.ContentTypeVideo, nil
	case mimeType == "application/pdf":
		return entity.ContentTypePDF, nil
	case strings.HasPrefix(mimeType, "image/"):
		return entity.ContentTypeImage, nil
	default:
		return "", ErrUnsupportedType
	}
}

// UploadContentUseCase handles uploading videos, PDFs, and images.
type UploadContentUseCase struct {
	contentRepo repository.ContentRepository
	lessonRepo  repository.LessonRepository
	storage     service.StorageService
	idGen       port.IDGenerator
}

// NewUploadContentUseCase creates a new UploadContentUseCase.
func NewUploadContentUseCase(
	contentRepo repository.ContentRepository,
	lessonRepo repository.LessonRepository,
	storage service.StorageService,
	idGen port.IDGenerator,
) *UploadContentUseCase {
	return &UploadContentUseCase{
		contentRepo: contentRepo,
		lessonRepo:  lessonRepo,
		storage:     storage,
		idGen:       idGen,
	}
}

// Execute uploads a file to private storage and records its metadata.
func (uc *UploadContentUseCase) Execute(ctx context.Context, req *dto.UploadContentRequest) (*dto.ContentResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, req.LessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	contentType, err := TypeFromMime(req.ContentType)
	if err != nil {
		return nil, err
	}

	id := uc.idGen.Generate()
	// Keys are namespaced per lesson and content ID so they never collide.
	storageKey := fmt.Sprintf("content/%s/%s%s", req.LessonID, id, strings.ToLower(path.Ext(req.FileName)))

	if err := uc.storage.Upload(ctx, storageKey, req.File, req.ContentType, req.Size); err != nil {
		return nil, err
	}

	now := time.Now()
	content := &entity.Content{
		ID:         id,
		LessonID:   req.LessonID,
		Type:       contentType,
		Title:      req.Title,
		StorageKey: storageKey,
		MimeType:   req.ContentType,
		Size:       req.Size,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := uc.contentRepo.Create(ctx, content); err != nil {
		// Best effort: don't leave an orphaned object behind.
		_ = uc.storage.Delete(ctx, storageKey)
		return nil, err
	}

	return &dto.ContentResponse{
		ID:        content.ID,
		LessonID:  content.LessonID,
		Type:      string(content.Type),
		Title:     content.Title,
		MimeType:  content.MimeType,
		Size:      content.Size,
		Duration:  content.Duration,
		Order:     content.Order,
		CreatedAt: content.CreatedAt,
	}, nil
}
