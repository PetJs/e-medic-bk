package quiz

import (
	"context"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// UpdateQuestionUseCase handles quiz question updates (prompt/order only —
// options aren't editable after creation in this slice).
type UpdateQuestionUseCase struct {
	questionRepo repository.QuizQuestionRepository
}

// NewUpdateQuestionUseCase creates a new UpdateQuestionUseCase.
func NewUpdateQuestionUseCase(questionRepo repository.QuizQuestionRepository) *UpdateQuestionUseCase {
	return &UpdateQuestionUseCase{questionRepo: questionRepo}
}

// Execute updates a quiz question's prompt and/or order.
func (uc *UpdateQuestionUseCase) Execute(ctx context.Context, id string, req *dto.QuizUpdateQuestionRequest) (*dto.QuizQuestionAdminResponse, error) {
	question, err := uc.questionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, ErrQuestionNotFound
	}

	if req.Prompt != nil {
		question.Prompt = *req.Prompt
	}
	if req.Order != nil {
		question.Order = *req.Order
	}
	question.UpdatedAt = time.Now()

	if err := uc.questionRepo.Update(ctx, question); err != nil {
		return nil, err
	}
	return ToQuestionAdminResponse(question), nil
}
