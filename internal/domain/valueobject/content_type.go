// Package valueobject contains immutable value objects with validation.
package valueobject

import "errors"

// ErrInvalidContentType is returned when a content type is invalid.
var ErrInvalidContentType = errors.New("invalid content type")

// ContentType represents the type of lesson content.
type ContentType string

const (
	ContentTypePDF   ContentType = "pdf"
	ContentTypeVideo ContentType = "video"
)

// NewContentType creates a new ContentType value object with validation.
func NewContentType(contentType string) (ContentType, error) {
	switch ContentType(contentType) {
	case ContentTypePDF, ContentTypeVideo:
		return ContentType(contentType), nil
	default:
		return "", ErrInvalidContentType
	}
}

// String returns the content type as a string.
func (c ContentType) String() string {
	return string(c)
}

// IsPDF returns true if the content type is PDF.
func (c ContentType) IsPDF() bool {
	return c == ContentTypePDF
}

// IsVideo returns true if the content type is video.
func (c ContentType) IsVideo() bool {
	return c == ContentTypeVideo
}
