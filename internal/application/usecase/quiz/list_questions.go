package quiz

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ListQuestionsUseCase lists a lesson's quiz questions for students (no is_correct).
type ListQuestionsUseCase struct {
	questionRepo repository.QuizQuestionRepository
	lessonRepo   repository.LessonRepository
}

// NewListQuestionsUseCase creates a new ListQuestionsUseCase.
func NewListQuestionsUseCase(questionRepo repository.QuizQuestionRepository, lessonRepo repository.LessonRepository) *ListQuestionsUseCase {
	return &ListQuestionsUseCase{questionRepo: questionRepo, lessonRepo: lessonRepo}
}

// Execute lists a lesson's quiz questions in student-facing shape.
func (uc *ListQuestionsUseCase) Execute(ctx context.Context, lessonID string) ([]*dto.QuizQuestionResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	questions, err := uc.questionRepo.ListByLesson(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	responses := make([]*dto.QuizQuestionResponse, 0, len(questions))
	for _, q := range questions {
		responses = append(responses, ToQuestionResponse(q))
	}
	return responses, nil
}

// ListQuestionsAdminUseCase lists a lesson's quiz questions for admins (includes is_correct).
type ListQuestionsAdminUseCase struct {
	questionRepo repository.QuizQuestionRepository
	lessonRepo   repository.LessonRepository
}

// NewListQuestionsAdminUseCase creates a new ListQuestionsAdminUseCase.
func NewListQuestionsAdminUseCase(questionRepo repository.QuizQuestionRepository, lessonRepo repository.LessonRepository) *ListQuestionsAdminUseCase {
	return &ListQuestionsAdminUseCase{questionRepo: questionRepo, lessonRepo: lessonRepo}
}

// Execute lists a lesson's quiz questions in admin-facing shape.
func (uc *ListQuestionsAdminUseCase) Execute(ctx context.Context, lessonID string) ([]*dto.QuizQuestionAdminResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	questions, err := uc.questionRepo.ListByLesson(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	responses := make([]*dto.QuizQuestionAdminResponse, 0, len(questions))
	for _, q := range questions {
		responses = append(responses, ToQuestionAdminResponse(q))
	}
	return responses, nil
}
