// Package progress contains progress tracking use cases.
package progress

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrLessonNotFound is returned when the lesson does not exist.
var ErrLessonNotFound = errors.New("lesson not found")

// ErrQuizIncomplete is returned when the caller explicitly asked to mark a
// lesson complete (or the video just ended) but the lesson has quiz
// questions the student hasn't answered yet.
var ErrQuizIncomplete = errors.New("quiz not yet completed")

func toProgressResponse(p *entity.Progress) *dto.ProgressResponse {
	return &dto.ProgressResponse{
		ID:           p.ID,
		UserID:       p.UserID,
		LessonID:     p.LessonID,
		ModuleID:     p.ModuleID,
		ProgressPct:  p.ProgressPct,
		LastPosition: p.LastPosition,
		IsCompleted:  p.IsCompleted,
		CompletedAt:  p.CompletedAt,
	}
}

// UpdateProgressUseCase handles updating lesson progress.
type UpdateProgressUseCase struct {
	progressRepo     repository.ProgressRepository
	lessonRepo       repository.LessonRepository
	quizQuestionRepo repository.QuizQuestionRepository
	quizAnswerRepo   repository.QuizAnswerRepository
	idGen            port.IDGenerator
}

// NewUpdateProgressUseCase creates a new UpdateProgressUseCase.
func NewUpdateProgressUseCase(
	progressRepo repository.ProgressRepository,
	lessonRepo repository.LessonRepository,
	quizQuestionRepo repository.QuizQuestionRepository,
	quizAnswerRepo repository.QuizAnswerRepository,
	idGen port.IDGenerator,
) *UpdateProgressUseCase {
	return &UpdateProgressUseCase{
		progressRepo:     progressRepo,
		lessonRepo:       lessonRepo,
		quizQuestionRepo: quizQuestionRepo,
		quizAnswerRepo:   quizAnswerRepo,
		idGen:            idGen,
	}
}

// Execute upserts the user's progress on a lesson. Progress never regresses:
// the stored percentage only grows and a completed lesson stays completed.
func (uc *UpdateProgressUseCase) Execute(ctx context.Context, userID string, req *dto.UpdateProgressRequest) (*dto.ProgressResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, req.LessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	now := time.Now()
	existing, err := uc.progressRepo.GetByUserAndLesson(ctx, userID, req.LessonID)
	if err != nil {
		return nil, err
	}

	p := existing
	if p == nil {
		p = &entity.Progress{
			ID:        uc.idGen.Generate(),
			UserID:    userID,
			LessonID:  req.LessonID,
			CreatedAt: now,
		}
	}

	pct := req.ProgressPct
	if pct < 0 {
		pct = 0
	}
	if pct > 100 {
		pct = 100
	}
	if pct > p.ProgressPct {
		p.ProgressPct = pct
	}
	if req.LastPosition > 0 {
		p.LastPosition = req.LastPosition
	}
	blockedByQuiz := false
	if (req.IsCompleted || p.ProgressPct >= 100) && !p.IsCompleted {
		total, err := uc.quizQuestionRepo.CountByLesson(ctx, req.LessonID)
		if err != nil {
			return nil, err
		}
		quizDone := true
		if total > 0 {
			answered, err := uc.quizAnswerRepo.CountAnsweredByUser(ctx, req.LessonID, userID)
			if err != nil {
				return nil, err
			}
			quizDone = answered >= total
		}

		if quizDone {
			p.IsCompleted = true
			p.ProgressPct = 100
			p.CompletedAt = &now
		} else if req.IsCompleted {
			// Explicit "mark complete" (or the video just ended) but the
			// lesson's quiz isn't done yet — block completion, but the
			// progress_pct/last_position updates above still get saved
			// below. An implicit pct>=100 with no explicit is_completed
			// just doesn't complete yet; no error, since that path is a
			// background progress ping, not a user action.
			blockedByQuiz = true
		}
	}
	p.UpdatedAt = now

	if err := uc.progressRepo.Upsert(ctx, p); err != nil {
		return nil, err
	}
	if blockedByQuiz {
		return nil, ErrQuizIncomplete
	}
	return toProgressResponse(p), nil
}
