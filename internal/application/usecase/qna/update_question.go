// Package qna contains Q&A system use cases.
package qna

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UpdateQuestionUseCase handles updating a question.
type UpdateQuestionUseCase struct{}

func NewUpdateQuestionUseCase() *UpdateQuestionUseCase { return &UpdateQuestionUseCase{} }

func (uc *UpdateQuestionUseCase) Execute(ctx context.Context, userID, questionID string, req *dto.UpdateQuestionRequest) (*dto.QuestionResponse, error) {
	return nil, nil
}
