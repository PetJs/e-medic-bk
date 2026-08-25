// Package quiz contains lesson quiz (post-lesson Q&A) use cases.
package quiz

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrLessonNotFound is returned when the parent lesson does not exist.
var ErrLessonNotFound = errors.New("lesson not found")

// ErrQuestionNotFound is returned when a question does not exist.
var ErrQuestionNotFound = errors.New("question not found")

// ErrInvalidQuestionType is returned when Type isn't a known question type.
// The validate:"..." DTO tags in this codebase are not actually enforced by
// gin's ShouldBindJSON (only binding:"..." is), so this is checked by hand —
// the same pattern content.TypeFromMime uses for Content.Type.
var ErrInvalidQuestionType = errors.New("invalid question type")

// ErrInvalidOptions is returned when a multiple_choice question's options
// don't have at least 2 entries with exactly one marked correct.
var ErrInvalidOptions = errors.New("multiple_choice questions need at least 2 options with exactly one marked correct")

// ToQuestionResponse maps a question to the student-facing shape (no is_correct).
func ToQuestionResponse(q *entity.QuizQuestion) *dto.QuizQuestionResponse {
	resp := &dto.QuizQuestionResponse{
		ID:       q.ID,
		LessonID: q.LessonID,
		Type:     string(q.Type),
		Prompt:   q.Prompt,
		Order:    q.Order,
	}
	for _, o := range q.Options {
		resp.Options = append(resp.Options, &dto.QuizOptionResponse{ID: o.ID, Text: o.Text, Order: o.Order})
	}
	return resp
}

// ToQuestionAdminResponse maps a question to the admin-facing shape (includes is_correct).
func ToQuestionAdminResponse(q *entity.QuizQuestion) *dto.QuizQuestionAdminResponse {
	resp := &dto.QuizQuestionAdminResponse{
		ID:       q.ID,
		LessonID: q.LessonID,
		Type:     string(q.Type),
		Prompt:   q.Prompt,
		Order:    q.Order,
	}
	for _, o := range q.Options {
		resp.Options = append(resp.Options, &dto.QuizOptionAdminResponse{ID: o.ID, Text: o.Text, IsCorrect: o.IsCorrect, Order: o.Order})
	}
	return resp
}

// CreateQuestionUseCase handles quiz question creation.
type CreateQuestionUseCase struct {
	questionRepo repository.QuizQuestionRepository
	lessonRepo   repository.LessonRepository
	idGen        port.IDGenerator
}

// NewCreateQuestionUseCase creates a new CreateQuestionUseCase.
func NewCreateQuestionUseCase(
	questionRepo repository.QuizQuestionRepository,
	lessonRepo repository.LessonRepository,
	idGen port.IDGenerator,
) *CreateQuestionUseCase {
	return &CreateQuestionUseCase{questionRepo: questionRepo, lessonRepo: lessonRepo, idGen: idGen}
}

// Execute creates a new quiz question (and, for multiple_choice, its options).
func (uc *CreateQuestionUseCase) Execute(ctx context.Context, lessonID string, req *dto.QuizCreateQuestionRequest) (*dto.QuizQuestionAdminResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	qType := entity.QuestionType(req.Type)
	if qType != entity.QuestionTypeMultipleChoice && qType != entity.QuestionTypeFreeText {
		return nil, ErrInvalidQuestionType
	}

	var options []*entity.QuizOption
	if qType == entity.QuestionTypeMultipleChoice {
		correctCount := 0
		for _, o := range req.Options {
			if o.IsCorrect {
				correctCount++
			}
		}
		if len(req.Options) < 2 || correctCount != 1 {
			return nil, ErrInvalidOptions
		}
	}

	now := time.Now()
	question := &entity.QuizQuestion{
		ID:        uc.idGen.Generate(),
		LessonID:  lessonID,
		Type:      qType,
		Prompt:    req.Prompt,
		Order:     req.Order,
		CreatedAt: now,
		UpdatedAt: now,
	}
	for i, o := range req.Options {
		options = append(options, &entity.QuizOption{
			ID:         uc.idGen.Generate(),
			QuestionID: question.ID,
			Text:       o.Text,
			IsCorrect:  o.IsCorrect,
			Order:      i,
		})
	}

	if err := uc.questionRepo.CreateWithOptions(ctx, question, options); err != nil {
		return nil, err
	}
	question.Options = options
	return ToQuestionAdminResponse(question), nil
}
