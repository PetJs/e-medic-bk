// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// UpdateProgressRequest represents a progress update request.
type UpdateProgressRequest struct {
	LessonID     string `json:"lesson_id" validate:"required"`
	ProgressPct  int    `json:"progress_pct" validate:"min=0,max=100"`
	LastPosition int    `json:"last_position,omitempty"` // for videos
	IsCompleted  bool   `json:"is_completed"`
}

// ProgressResponse represents progress in API responses.
type ProgressResponse struct {
	ID           string     `json:"id"`
	UserID       string     `json:"user_id"`
	LessonID     string     `json:"lesson_id"`
	ProgressPct  int        `json:"progress_pct"`
	LastPosition int        `json:"last_position"`
	IsCompleted  bool       `json:"is_completed"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
}

// CourseProgressResponse represents course progress stats.
type CourseProgressResponse struct {
	CourseID         string  `json:"course_id"`
	CompletedLessons int     `json:"completed_lessons"`
	TotalLessons     int     `json:"total_lessons"`
	ProgressPct      float64 `json:"progress_pct"`
}
