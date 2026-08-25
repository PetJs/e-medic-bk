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

// ErrOptionNotFound is returned when the selected option doesn't belong to
// the target question.
var ErrOptionNotFound = errors.New("option not found")

// ErrAnswerShapeMismatch is returned when the submission doesn't match the
// question's type (e.g. free text on a multiple_choice question).
var ErrAnswerShapeMismatch = errors.New("answer does not match question type")

func toAnswerResponse(a *entity.QuizAnswer) *dto.QuizAnswerResponse {
	resp := &dto.QuizAnswerResponse{
		ID:               a.ID,
		QuestionID:       a.QuestionID,
		UserID:           a.UserID,
		SelectedOptionID: a.SelectedOptionID,
		FreeTextBody:     a.FreeTextBody,
		IsCorrect:        a.IsCorrect,
		CreatedAt:        a.CreatedAt,
	}
	if a.Author != nil {
		resp.Author = &dto.UserResponse{
			ID:        a.Author.ID,
			Email:     a.Author.Email,
			FirstName: a.Author.FirstName,
			LastName:  a.Author.LastName,
			Role:      a.Author.Role,
			CreatedAt: a.Author.CreatedAt,
		}
	}
	return resp
}

// SubmitAnswerUseCase handles a student's answer submission to one question.
type SubmitAnswerUseCase struct {
	answerRepo   repository.QuizAnswerRepository
	questionRepo repository.QuizQuestionRepository
	idGen        port.IDGenerator
}

// NewSubmitAnswerUseCase creates a new SubmitAnswerUseCase.
func NewSubmitAnswerUseCase(
	answerRepo repository.QuizAnswerRepository,
	questionRepo repository.QuizQuestionRepository,
	idGen port.IDGenerator,
) *SubmitAnswerUseCase {
	return &SubmitAnswerUseCase{answerRepo: answerRepo, questionRepo: questionRepo, idGen: idGen}
}

// Execute submits (once — locked in by a DB unique constraint) a student's
// answer to a question. multiple_choice is graded immediately; free_text is
// stored ungraded for admin review.
func (uc *SubmitAnswerUseCase) Execute(ctx context.Context, questionID, userID string, req *dto.QuizSubmitAnswerRequest) (*dto.QuizAnswerResponse, error) {
	question, err := uc.questionRepo.GetByID(ctx, questionID)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, ErrQuestionNotFound
	}

	answer := &entity.QuizAnswer{
		ID:         uc.idGen.Generate(),
		QuestionID: questionID,
		UserID:     userID,
		CreatedAt:  time.Now(),
	}

	switch question.Type {
	case entity.QuestionTypeMultipleChoice:
		if req.SelectedOptionID == nil {
			return nil, ErrAnswerShapeMismatch
		}
		var selected *entity.QuizOption
		for _, o := range question.Options {
			if o.ID == *req.SelectedOptionID {
				selected = o
				break
			}
		}
		if selected == nil {
			return nil, ErrOptionNotFound
		}
		correct := selected.IsCorrect
		answer.SelectedOptionID = req.SelectedOptionID
		answer.IsCorrect = &correct
	case entity.QuestionTypeFreeText:
		if req.FreeTextBody == nil || *req.FreeTextBody == "" {
			return nil, ErrAnswerShapeMismatch
		}
		answer.FreeTextBody = req.FreeTextBody
	}

	if err := uc.answerRepo.Create(ctx, answer); err != nil {
		return nil, err
	}
	return toAnswerResponse(answer), nil
}
