// Package course contains course management use cases.
package course

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrCourseNotFound is returned when a course does not exist.
var ErrCourseNotFound = errors.New("course not found")

var slugPattern = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converts a title into a URL-friendly slug.
func Slugify(title string) string {
	slug := strings.ToLower(strings.TrimSpace(title))
	slug = slugPattern.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

func toCourseResponse(c *entity.Course) *dto.CourseResponse {
	return &dto.CourseResponse{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		Slug:        c.Slug,
		CoverImage:  c.CoverImage,
		AuthorID:    c.AuthorID,
		IsPublished: c.IsPublished,
		CreatedAt:   c.CreatedAt,
	}
}

// CreateCourseUseCase handles course creation.
type CreateCourseUseCase struct {
	courseRepo repository.CourseRepository
	idGen      port.IDGenerator
}

// NewCreateCourseUseCase creates a new CreateCourseUseCase.
func NewCreateCourseUseCase(courseRepo repository.CourseRepository, idGen port.IDGenerator) *CreateCourseUseCase {
	return &CreateCourseUseCase{courseRepo: courseRepo, idGen: idGen}
}

// Execute creates a new course.
func (uc *CreateCourseUseCase) Execute(ctx context.Context, authorID string, req *dto.CreateCourseRequest) (*dto.CourseResponse, error) {
	slug := Slugify(req.Title)
	// Ensure slug uniqueness by suffixing a counter if taken
	base := slug
	for i := 2; ; i++ {
		existing, err := uc.courseRepo.GetBySlug(ctx, slug)
		if err != nil {
			return nil, err
		}
		if existing == nil {
			break
		}
		slug = fmt.Sprintf("%s-%d", base, i)
	}

	now := time.Now()
	course := &entity.Course{
		ID:          uc.idGen.Generate(),
		Title:       req.Title,
		Description: req.Description,
		Slug:        slug,
		CoverImage:  req.CoverImage,
		AuthorID:    authorID,
		IsPublished: true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := uc.courseRepo.Create(ctx, course); err != nil {
		return nil, err
	}
	return toCourseResponse(course), nil
}
