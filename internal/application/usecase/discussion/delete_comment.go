package discussion

import (
	"context"
	"errors"

	"emedic-bk/internal/domain/repository"
)

// ErrCommentNotFound is returned when a comment does not exist, or the
// caller is not allowed to act on it (kept indistinguishable from "not
// found" so ownership checks don't leak existence of other users' comments).
var ErrCommentNotFound = errors.New("comment not found")

// DeleteCommentUseCase handles discussion comment deletion.
type DeleteCommentUseCase struct {
	commentRepo repository.DiscussionCommentRepository
}

// NewDeleteCommentUseCase creates a new DeleteCommentUseCase.
func NewDeleteCommentUseCase(commentRepo repository.DiscussionCommentRepository) *DeleteCommentUseCase {
	return &DeleteCommentUseCase{commentRepo: commentRepo}
}

// Execute deletes a comment (and, via DB cascade, every reply beneath it).
// Only the comment's author or an admin may delete it.
func (uc *DeleteCommentUseCase) Execute(ctx context.Context, userID, role, commentID string) error {
	comment, err := uc.commentRepo.GetByID(ctx, commentID)
	if err != nil {
		return err
	}
	if comment == nil {
		return ErrCommentNotFound
	}
	if role != "admin" && comment.UserID != userID {
		return ErrCommentNotFound
	}
	return uc.commentRepo.Delete(ctx, commentID)
}
