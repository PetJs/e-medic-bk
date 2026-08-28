// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// LessonNoteRepository defines the interface for personal lesson-note data access.
type LessonNoteRepository interface {
	Create(ctx context.Context, note *entity.LessonNote) error
	GetByID(ctx context.Context, id string) (*entity.LessonNote, error)
	Delete(ctx context.Context, id string) error
	// ListByUserAndLesson returns a student's own notes on a lesson, ordered
	// by the video moment they're tied to.
	ListByUserAndLesson(ctx context.Context, userID, lessonID string) ([]*entity.LessonNote, error)
}
