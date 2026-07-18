// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// CourseRepository implements repository.CourseRepository.
type CourseRepository struct {
	db *DB
}

func NewCourseRepository(db *DB) repository.CourseRepository {
	return &CourseRepository{db: db}
}

// description and cover_image are nullable columns — coalesce for plain string scans.
const courseColumns = `id, title, COALESCE(description, ''), slug, COALESCE(cover_image, ''), author_id, is_published, created_at, updated_at`

func (r *CourseRepository) Create(ctx context.Context, course *entity.Course) error {
	query := `
		INSERT INTO courses (id, title, description, slug, cover_image, author_id, is_published, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		course.ID,
		course.Title,
		course.Description,
		course.Slug,
		course.CoverImage,
		course.AuthorID,
		course.IsPublished,
		course.CreatedAt,
		course.UpdatedAt,
	)
	return err
}

func (r *CourseRepository) GetByID(ctx context.Context, id string) (*entity.Course, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+courseColumns+` FROM courses WHERE id = $1`, id)
	return r.scanCourse(row)
}

func (r *CourseRepository) GetBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+courseColumns+` FROM courses WHERE slug = $1`, slug)
	return r.scanCourse(row)
}

func (r *CourseRepository) Update(ctx context.Context, course *entity.Course) error {
	query := `
		UPDATE courses
		SET title = $2, description = $3, slug = $4, cover_image = $5, is_published = $6, updated_at = $7
		WHERE id = $1
	`
	_, err := r.db.Pool.Exec(ctx, query,
		course.ID,
		course.Title,
		course.Description,
		course.Slug,
		course.CoverImage,
		course.IsPublished,
		time.Now(),
	)
	return err
}

func (r *CourseRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM courses WHERE id = $1`, id)
	return err
}

func (r *CourseRepository) List(ctx context.Context, limit, offset int) ([]*entity.Course, error) {
	query := `SELECT ` + courseColumns + ` FROM courses ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanCourses(rows)
}

func (r *CourseRepository) ListByAuthor(ctx context.Context, authorID string, limit, offset int) ([]*entity.Course, error) {
	query := `SELECT ` + courseColumns + ` FROM courses WHERE author_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Pool.Query(ctx, query, authorID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanCourses(rows)
}

func (r *CourseRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM courses`).Scan(&count)
	return count, err
}

func (r *CourseRepository) scanCourse(row pgx.Row) (*entity.Course, error) {
	course := &entity.Course{}
	err := row.Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.Slug,
		&course.CoverImage,
		&course.AuthorID,
		&course.IsPublished,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return course, nil
}

func (r *CourseRepository) scanCourses(rows pgx.Rows) ([]*entity.Course, error) {
	var courses []*entity.Course
	for rows.Next() {
		course := &entity.Course{}
		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.Slug,
			&course.CoverImage,
			&course.AuthorID,
			&course.IsPublished,
			&course.CreatedAt,
			&course.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, rows.Err()
}
