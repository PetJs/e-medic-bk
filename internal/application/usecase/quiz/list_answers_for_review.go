package quiz

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ListAnswersForReviewUseCase lists every student's answer to one question —
// admin-only, used to review free_text submissions (multiple_choice answers
// are already auto-graded, but are included too for completeness).
type ListAnswersForReviewUseCase struct {
	answerRepo   repository.QuizAnswerRepository
	questionRepo repository.QuizQuestionRepository
}

// NewListAnswersForReviewUseCase creates a new ListAnswersForReviewUseCase.
func NewListAnswersForReviewUseCase(answerRepo repository.QuizAnswerRepository, questionRepo repository.QuizQuestionRepository) *ListAnswersForReviewUseCase {
	return &ListAnswersForReviewUseCase{answerRepo: answerRepo, questionRepo: questionRepo}
}

// Execute returns every submitted answer to one question.
func (uc *ListAnswersForReviewUseCase) Execute(ctx context.Context, questionID string) (*dto.QuizListAnswersResponse, error) {
	question, err := uc.questionRepo.GetByID(ctx, questionID)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, ErrQuestionNotFound
	}

	answers, err := uc.answerRepo.ListByQuestion(ctx, questionID)
	if err != nil {
		return nil, err
	}
	responses := make([]*dto.QuizAnswerResponse, 0, len(answers))
	for _, a := range answers {
		responses = append(responses, toAnswerResponse(a))
	}
	return &dto.QuizListAnswersResponse{Answers: responses}, nil
}
