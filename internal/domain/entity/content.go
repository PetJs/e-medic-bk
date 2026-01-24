// Package entity contains the core domain entities.
package entity

import "time"

// ContentType represents the type of content.
type ContentType string

const (
	ContentTypePDF   ContentType = "pdf"
	ContentTypeVideo ContentType = "video"
)

// Content represents a piece of content (PDF or video) attached to a lesson.
type Content struct {
	ID          string
	LessonID    string
	Type        ContentType
	Title       string
	StorageKey  string // S3 object key
	MimeType    string
	Size        int64 // in bytes
	Duration    int   // for videos, in seconds
	Order       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
