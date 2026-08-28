package notes

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ListNotesUseCase handles listing a student's own notes on a lesson.
type ListNotesUseCase struct {
	noteRepo repository.LessonNoteRepository
}

// NewListNotesUseCase creates a new ListNotesUseCase.
func NewListNotesUseCase(noteRepo repository.LessonNoteRepository) *ListNotesUseCase {
	return &ListNotesUseCase{noteRepo: noteRepo}
}

// Execute returns the caller's own notes on a lesson, ordered by video position.
func (uc *ListNotesUseCase) Execute(ctx context.Context, userID, lessonID string) (*dto.ListNotesResponse, error) {
	notes, err := uc.noteRepo.ListByUserAndLesson(ctx, userID, lessonID)
	if err != nil {
		return nil, err
	}
	responses := make([]*dto.NoteResponse, 0, len(notes))
	for _, n := range notes {
		responses = append(responses, toNoteResponse(n))
	}
	return &dto.ListNotesResponse{Notes: responses}, nil
}
