// Package dto contains Data Transfer Objects for use cases.
package dto

import (
	"io"
	"time"
)

// UploadContentRequest represents a content upload request.
type UploadContentRequest struct {
	LessonID    string    `json:"lesson_id" validate:"required"`
	Type        string    `json:"type" validate:"required,oneof=pdf video"`
	Title       string    `json:"title" validate:"required"`
	File        io.Reader `json:"-"`
	FileName    string    `json:"-"`
	ContentType string    `json:"-"`
	Size        int64     `json:"-"`
}

// ContentResponse represents content in API responses.
type ContentResponse struct {
	ID        string    `json:"id"`
	LessonID  string    `json:"lesson_id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	MimeType  string    `json:"mime_type"`
	Size      int64     `json:"size"`
	Duration  int       `json:"duration,omitempty"` // for videos
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
}

// ContentURLResponse represents a signed URL for content access.
type ContentURLResponse struct {
	URL       string    `json:"url"`
	ExpiresAt time.Time `json:"expires_at"`
}
