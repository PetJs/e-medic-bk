// Package qna contains Q&A system use cases.
package qna

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// CreateQuestionUseCase handles creating a question.
type CreateQuestionUseCase struct{}

func NewCreateQuestionUseCase() *CreateQuestionUseCase { return &CreateQuestionUseCase{} }

func (uc *CreateQuestionUseCase) Execute(ctx context.Context, userID string, req *dto.CreateQuestionRequest) (*dto.QuestionResponse, error) {
	return nil, nil
}
