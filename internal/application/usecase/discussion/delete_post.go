package discussion

import (
	"context"

	"emedic-bk/internal/domain/repository"
)

// DeletePostUseCase handles discussion post deletion.
type DeletePostUseCase struct {
	postRepo repository.DiscussionPostRepository
}

// NewDeletePostUseCase creates a new DeletePostUseCase.
func NewDeletePostUseCase(postRepo repository.DiscussionPostRepository) *DeletePostUseCase {
	return &DeletePostUseCase{postRepo: postRepo}
}

// Execute deletes a post. Only the post's author or an admin may delete it;
// anyone else gets ErrPostNotFound so existence isn't leaked (comments cascade
// via the DB foreign key).
func (uc *DeletePostUseCase) Execute(ctx context.Context, userID, role, postID string) error {
	post, err := uc.postRepo.GetByID(ctx, postID)
	if err != nil {
		return err
	}
	if post == nil {
		return ErrPostNotFound
	}
	if role != "admin" && post.UserID != userID {
		return ErrPostNotFound
	}
	return uc.postRepo.Delete(ctx, postID)
}
