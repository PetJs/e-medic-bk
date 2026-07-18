// Package module contains module management use cases.
package module

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrModuleNotFound is returned when a module does not exist.
var ErrModuleNotFound = errors.New("module not found")

// ErrCourseNotFound is returned when the parent course does not exist.
var ErrCourseNotFound = errors.New("course not found")

func toModuleResponse(m *entity.Module) *dto.ModuleResponse {
	return &dto.ModuleResponse{
		ID:            m.ID,
		CourseID:      m.CourseID,
		Title:         m.Title,
		Description:   m.Description,
		Order:         m.Order,
		IsPremium:     m.IsPremium,
		LessonCount:   m.LessonCount,
		TotalDuration: m.TotalDuration,
		CreatedAt:     m.CreatedAt,
	}
}

// CreateModuleUseCase handles module creation.
type CreateModuleUseCase struct {
	moduleRepo repository.ModuleRepository
	courseRepo repository.CourseRepository
	idGen      port.IDGenerator
}

// NewCreateModuleUseCase creates a new CreateModuleUseCase.
func NewCreateModuleUseCase(
	moduleRepo repository.ModuleRepository,
	courseRepo repository.CourseRepository,
	idGen port.IDGenerator,
) *CreateModuleUseCase {
	return &CreateModuleUseCase{moduleRepo: moduleRepo, courseRepo: courseRepo, idGen: idGen}
}

// Execute creates a new module.
func (uc *CreateModuleUseCase) Execute(ctx context.Context, req *dto.CreateModuleRequest) (*dto.ModuleResponse, error) {
	course, err := uc.courseRepo.GetByID(ctx, req.CourseID)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, ErrCourseNotFound
	}

	now := time.Now()
	module := &entity.Module{
		ID:          uc.idGen.Generate(),
		CourseID:    req.CourseID,
		Title:       req.Title,
		Description: req.Description,
		Order:       req.Order,
		IsPremium:   req.IsPremium,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := uc.moduleRepo.Create(ctx, module); err != nil {
		return nil, err
	}
	return toModuleResponse(module), nil
}
