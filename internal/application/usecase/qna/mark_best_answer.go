// Package qna contains Q&A system use cases.
package qna

import "context"

// MarkBestAnswerUseCase handles marking an answer as best.
type MarkBestAnswerUseCase struct{}

func NewMarkBestAnswerUseCase() *MarkBestAnswerUseCase { return &MarkBestAnswerUseCase{} }

func (uc *MarkBestAnswerUseCase) Execute(ctx context.Context, userID, questionID, answerID string) error {
	// TODO: Verify user owns the question
	// TODO: Clear existing best answer
	// TODO: Mark new best answer
	return nil
}
