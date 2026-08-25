package quiz

import (
	"context"

	"emedic-bk/internal/domain/repository"
)

// DeleteQuestionUseCase handles quiz question deletion (admin).
type DeleteQuestionUseCase struct {
	questionRepo repository.QuizQuestionRepository
}

// NewDeleteQuestionUseCase creates a new DeleteQuestionUseCase.
func NewDeleteQuestionUseCase(questionRepo repository.QuizQuestionRepository) *DeleteQuestionUseCase {
	return &DeleteQuestionUseCase{questionRepo: questionRepo}
}

// Execute deletes a quiz question (its options and any submitted answers
// cascade via the DB foreign keys).
func (uc *DeleteQuestionUseCase) Execute(ctx context.Context, id string) error {
	question, err := uc.questionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if question == nil {
		return ErrQuestionNotFound
	}
	return uc.questionRepo.Delete(ctx, id)
}
