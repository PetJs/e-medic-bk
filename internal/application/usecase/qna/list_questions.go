// Package qna contains Q&A system use cases.
package qna

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ListQuestionsUseCase handles listing questions by lesson.
type ListQuestionsUseCase struct{}

func NewListQuestionsUseCase() *ListQuestionsUseCase { return &ListQuestionsUseCase{} }

func (uc *ListQuestionsUseCase) Execute(ctx context.Context, lessonID string, req *dto.ListQuestionsRequest) (*dto.ListQuestionsResponse, error) {
	return nil, nil
}
