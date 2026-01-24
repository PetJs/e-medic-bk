// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreateLessonRequest represents a lesson creation request.
type CreateLessonRequest struct {
	ModuleID    string `json:"module_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	Duration    int    `json:"duration"` // in minutes
}

// UpdateLessonRequest represents a lesson update request.
type UpdateLessonRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Order       *int    `json:"order,omitempty"`
	Duration    *int    `json:"duration,omitempty"`
}

// LessonResponse represents a lesson in API responses.
type LessonResponse struct {
	ID          string    `json:"id"`
	ModuleID    string    `json:"module_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
}

// LessonDetailResponse represents detailed lesson info with content.
type LessonDetailResponse struct {
	LessonResponse
	Contents []*ContentResponse `json:"contents"`
	Progress *ProgressResponse  `json:"progress,omitempty"`
}
