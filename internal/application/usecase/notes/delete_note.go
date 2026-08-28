package notes

import (
	"context"

	"emedic-bk/internal/domain/repository"
)

// DeleteNoteUseCase handles personal note deletion.
type DeleteNoteUseCase struct {
	noteRepo repository.LessonNoteRepository
}

// NewDeleteNoteUseCase creates a new DeleteNoteUseCase.
func NewDeleteNoteUseCase(noteRepo repository.LessonNoteRepository) *DeleteNoteUseCase {
	return &DeleteNoteUseCase{noteRepo: noteRepo}
}

// Execute deletes a note. Only its own author may delete it — deliberately
// no admin override, since these are private notes, not moderated content.
func (uc *DeleteNoteUseCase) Execute(ctx context.Context, userID, noteID string) error {
	note, err := uc.noteRepo.GetByID(ctx, noteID)
	if err != nil {
		return err
	}
	if note == nil || note.UserID != userID {
		return ErrNoteNotFound
	}
	return uc.noteRepo.Delete(ctx, noteID)
}
