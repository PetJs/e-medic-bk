// Package module contains module management use cases.
package module

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// ErrModuleNotFound is returned when a module does not exist.
var ErrModuleNotFound = errors.New("module not found")

// ErrCourseNotFound is returned when the parent course does not exist.
var ErrCourseNotFound = errors.New("course not found")

// ToModuleResponse maps a module entity to its API response, resolving a
// signed URL for the cover image when one has finished generating.
func ToModuleResponse(ctx context.Context, m *entity.Module, storageSvc service.StorageService) *dto.ModuleResponse {
	resp := &dto.ModuleResponse{
		ID:               m.ID,
		CourseID:         m.CourseID,
		Title:            m.Title,
		Description:      m.Description,
		Order:            m.Order,
		IsPremium:        m.IsPremium,
		CoverImageStatus: m.CoverImageStatus,
		LessonCount:      m.LessonCount,
		TotalDuration:    m.TotalDuration,
		CreatedAt:        m.CreatedAt,
	}
	if m.CoverImageKey != "" && m.CoverImageStatus == "ready" && storageSvc != nil {
		url, err := storageSvc.GetSignedURL(ctx, m.CoverImageKey, time.Hour)
		if err != nil {
			slog.Error("failed to sign module cover URL", "module_id", m.ID, "error", err)
		} else {
			resp.CoverImageURL = url
		}
	}
	return resp
}

// CreateModuleUseCase handles module creation.
type CreateModuleUseCase struct {
	moduleRepo repository.ModuleRepository
	courseRepo repository.CourseRepository
	idGen      port.IDGenerator
	storageSvc service.StorageService
	coverGen   *ModuleCoverGenerator
}

// NewCreateModuleUseCase creates a new CreateModuleUseCase.
func NewCreateModuleUseCase(
	moduleRepo repository.ModuleRepository,
	courseRepo repository.CourseRepository,
	idGen port.IDGenerator,
	storageSvc service.StorageService,
	coverGen *ModuleCoverGenerator,
) *CreateModuleUseCase {
	return &CreateModuleUseCase{
		moduleRepo: moduleRepo,
		courseRepo: courseRepo,
		idGen:      idGen,
		storageSvc: storageSvc,
		coverGen:   coverGen,
	}
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
	uc.coverGen.Trigger(module.ID, module.Title, module.Description)
	return ToModuleResponse(ctx, module, uc.storageSvc), nil
}
