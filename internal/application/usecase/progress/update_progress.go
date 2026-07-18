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
	progressRepo repository.ProgressRepository
	lessonRepo   repository.LessonRepository
	idGen        port.IDGenerator
}

// NewUpdateProgressUseCase creates a new UpdateProgressUseCase.
func NewUpdateProgressUseCase(
	progressRepo repository.ProgressRepository,
	lessonRepo repository.LessonRepository,
	idGen port.IDGenerator,
) *UpdateProgressUseCase {
	return &UpdateProgressUseCase{progressRepo: progressRepo, lessonRepo: lessonRepo, idGen: idGen}
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
	if (req.IsCompleted || p.ProgressPct >= 100) && !p.IsCompleted {
		p.IsCompleted = true
		p.ProgressPct = 100
		p.CompletedAt = &now
	}
	p.UpdatedAt = now

	if err := uc.progressRepo.Upsert(ctx, p); err != nil {
		return nil, err
	}
	return toProgressResponse(p), nil
}
