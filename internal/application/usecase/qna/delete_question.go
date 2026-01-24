// Package qna contains Q&A system use cases.
package qna

import "context"

// DeleteQuestionUseCase handles deleting a question.
type DeleteQuestionUseCase struct{}

func NewDeleteQuestionUseCase() *DeleteQuestionUseCase { return &DeleteQuestionUseCase{} }

func (uc *DeleteQuestionUseCase) Execute(ctx context.Context, userID, questionID string) error {
	return nil
}
