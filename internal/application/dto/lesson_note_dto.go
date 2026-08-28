// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreateNoteRequest represents a personal-note creation request.
// The lesson is identified by the URL path, not the body.
type CreateNoteRequest struct {
	Body          string `json:"body" binding:"required"`
	VideoPosition int    `json:"video_position"`
}

// NoteResponse represents a personal lesson note in API responses.
type NoteResponse struct {
	ID            string    `json:"id"`
	LessonID      string    `json:"lesson_id"`
	Body          string    `json:"body"`
	VideoPosition int       `json:"video_position"`
	CreatedAt     time.Time `json:"created_at"`
}

// ListNotesResponse represents a student's notes on a lesson.
type ListNotesResponse struct {
	Notes []*NoteResponse `json:"notes"`
}
