// Package qna contains Q&A system use cases.
package qna

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UpdateAnswerUseCase handles updating an answer.
type UpdateAnswerUseCase struct{}

func NewUpdateAnswerUseCase() *UpdateAnswerUseCase { return &UpdateAnswerUseCase{} }

func (uc *UpdateAnswerUseCase) Execute(ctx context.Context, userID, answerID string, req *dto.UpdateAnswerRequest) (*dto.AnswerResponse, error) {
	return nil, nil
}
