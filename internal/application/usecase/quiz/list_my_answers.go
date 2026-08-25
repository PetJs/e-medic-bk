package quiz

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ListMyAnswersUseCase lists the calling student's own answers for a lesson.
type ListMyAnswersUseCase struct {
	answerRepo repository.QuizAnswerRepository
	lessonRepo repository.LessonRepository
}

// NewListMyAnswersUseCase creates a new ListMyAnswersUseCase.
func NewListMyAnswersUseCase(answerRepo repository.QuizAnswerRepository, lessonRepo repository.LessonRepository) *ListMyAnswersUseCase {
	return &ListMyAnswersUseCase{answerRepo: answerRepo, lessonRepo: lessonRepo}
}

// Execute returns the calling student's own answers for a lesson's quiz —
// used both to render "already answered" state and to know which questions
// remain.
func (uc *ListMyAnswersUseCase) Execute(ctx context.Context, lessonID, userID string) (*dto.QuizListAnswersResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	answers, err := uc.answerRepo.ListByUserAndLesson(ctx, userID, lessonID)
	if err != nil {
		return nil, err
	}
	responses := make([]*dto.QuizAnswerResponse, 0, len(answers))
	for _, a := range answers {
		responses = append(responses, toAnswerResponse(a))
	}
	return &dto.QuizListAnswersResponse{Answers: responses}, nil
}
