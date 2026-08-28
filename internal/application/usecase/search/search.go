// Package search contains the site-search use case.
package search

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// resultsPerCategory caps how many hits come back per entity type — this
// feeds a quick-search popup, not a paginated results page.
const resultsPerCategory = 5

// SearchUseCase searches courses, modules, and lessons by title/description.
type SearchUseCase struct {
	courseRepo repository.CourseRepository
	moduleRepo repository.ModuleRepository
	lessonRepo repository.LessonRepository
}

// NewSearchUseCase creates a new SearchUseCase.
func NewSearchUseCase(
	courseRepo repository.CourseRepository,
	moduleRepo repository.ModuleRepository,
	lessonRepo repository.LessonRepository,
) *SearchUseCase {
	return &SearchUseCase{courseRepo: courseRepo, moduleRepo: moduleRepo, lessonRepo: lessonRepo}
}

// Execute runs the same query against all three entity types in parallel-ish
// sequence (no goroutines — matches this codebase's existing sequential
// usecase style) and returns up to resultsPerCategory hits each.
func (uc *SearchUseCase) Execute(ctx context.Context, query string) (*dto.SearchResponse, error) {
	if query == "" {
		return &dto.SearchResponse{
			Courses: []*dto.SearchResultItem{},
			Modules: []*dto.SearchResultItem{},
			Lessons: []*dto.SearchResultItem{},
		}, nil
	}

	courses, err := uc.courseRepo.Search(ctx, query, resultsPerCategory)
	if err != nil {
		return nil, err
	}
	modules, err := uc.moduleRepo.Search(ctx, query, resultsPerCategory)
	if err != nil {
		return nil, err
	}
	lessons, err := uc.lessonRepo.Search(ctx, query, resultsPerCategory)
	if err != nil {
		return nil, err
	}

	resp := &dto.SearchResponse{
		Courses: make([]*dto.SearchResultItem, 0, len(courses)),
		Modules: make([]*dto.SearchResultItem, 0, len(modules)),
		Lessons: make([]*dto.SearchResultItem, 0, len(lessons)),
	}
	for _, c := range courses {
		resp.Courses = append(resp.Courses, toResultItem(c.ID, c.Title, c.Description))
	}
	for _, m := range modules {
		resp.Modules = append(resp.Modules, toResultItem(m.ID, m.Title, m.Description))
	}
	for _, l := range lessons {
		resp.Lessons = append(resp.Lessons, toResultItem(l.ID, l.Title, l.Description))
	}
	return resp, nil
}

func toResultItem(id, title, description string) *dto.SearchResultItem {
	return &dto.SearchResultItem{ID: id, Title: title, Description: description}
}
