// Package qna contains Q&A system use cases.
package qna

import "context"

// DeleteAnswerUseCase handles deleting an answer.
type DeleteAnswerUseCase struct{}

func NewDeleteAnswerUseCase() *DeleteAnswerUseCase { return &DeleteAnswerUseCase{} }

func (uc *DeleteAnswerUseCase) Execute(ctx context.Context, userID, answerID string) error {
	return nil
}
