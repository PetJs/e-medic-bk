// Package notes contains personal (private, per-student) lesson-note use cases.
package notes

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrLessonNotFound is returned when the target lesson does not exist.
var ErrLessonNotFound = errors.New("lesson not found")

// ErrNoteNotFound is returned when a note does not exist, or the caller
// isn't its owner — kept indistinguishable from "not found" so ownership
// checks don't leak existence of another student's private notes. Unlike
// discussion posts, there is deliberately NO admin bypass here: notes are
// private, and an admin has no legitimate reason to read or delete them.
var ErrNoteNotFound = errors.New("note not found")

func toNoteResponse(n *entity.LessonNote) *dto.NoteResponse {
	return &dto.NoteResponse{
		ID:            n.ID,
		LessonID:      n.LessonID,
		Body:          n.Body,
		VideoPosition: n.VideoPosition,
		CreatedAt:     n.CreatedAt,
	}
}

// CreateNoteUseCase handles personal note creation.
type CreateNoteUseCase struct {
	noteRepo   repository.LessonNoteRepository
	lessonRepo repository.LessonRepository
	idGen      port.IDGenerator
}

// NewCreateNoteUseCase creates a new CreateNoteUseCase.
func NewCreateNoteUseCase(
	noteRepo repository.LessonNoteRepository,
	lessonRepo repository.LessonRepository,
	idGen port.IDGenerator,
) *CreateNoteUseCase {
	return &CreateNoteUseCase{noteRepo: noteRepo, lessonRepo: lessonRepo, idGen: idGen}
}

// Execute creates a new timestamped note on a lesson.
func (uc *CreateNoteUseCase) Execute(ctx context.Context, lessonID, userID string, req *dto.CreateNoteRequest) (*dto.NoteResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	position := req.VideoPosition
	if position < 0 {
		position = 0
	}

	now := time.Now()
	note := &entity.LessonNote{
		ID:            uc.idGen.Generate(),
		UserID:        userID,
		LessonID:      lessonID,
		Body:          req.Body,
		VideoPosition: position,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	if err := uc.noteRepo.Create(ctx, note); err != nil {
		return nil, err
	}
	return toNoteResponse(note), nil
}
