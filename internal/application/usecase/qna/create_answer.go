// Package qna contains Q&A system use cases.
package qna

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// CreateAnswerUseCase handles creating an answer.
type CreateAnswerUseCase struct{}

func NewCreateAnswerUseCase() *CreateAnswerUseCase { return &CreateAnswerUseCase{} }

func (uc *CreateAnswerUseCase) Execute(ctx context.Context, userID string, req *dto.CreateAnswerRequest) (*dto.AnswerResponse, error) {
	return nil, nil
}
