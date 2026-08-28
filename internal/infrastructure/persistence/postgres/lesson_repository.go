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

type LessonRepository struct{ db *DB }

func NewLessonRepository(db *DB) repository.LessonRepository { return &LessonRepository{db: db} }

const lessonColumns = `id, module_id, title, COALESCE(description, ''), "order", duration, created_at, updated_at`

func (r *LessonRepository) Create(ctx context.Context, lesson *entity.Lesson) error {
	query := `
		INSERT INTO lessons (id, module_id, title, description, "order", duration, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		lesson.ID,
		lesson.ModuleID,
		lesson.Title,
		lesson.Description,
		lesson.Order,
		lesson.Duration,
		lesson.CreatedAt,
		lesson.UpdatedAt,
	)
	return err
}

func (r *LessonRepository) GetByID(ctx context.Context, id string) (*entity.Lesson, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+lessonColumns+` FROM lessons WHERE id = $1`, id)
	return r.scanLesson(row)
}

func (r *LessonRepository) Update(ctx context.Context, lesson *entity.Lesson) error {
	query := `
		UPDATE lessons
		SET title = $2, description = $3, "order" = $4, duration = $5, updated_at = $6
		WHERE id = $1
	`
	_, err := r.db.Pool.Exec(ctx, query,
		lesson.ID,
		lesson.Title,
		lesson.Description,
		lesson.Order,
		lesson.Duration,
		time.Now(),
	)
	return err
}

func (r *LessonRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM lessons WHERE id = $1`, id)
	return err
}

func (r *LessonRepository) ListByModule(ctx context.Context, moduleID string) ([]*entity.Lesson, error) {
	query := `SELECT ` + lessonColumns + ` FROM lessons WHERE module_id = $1 ORDER BY "order", created_at`
	rows, err := r.db.Pool.Query(ctx, query, moduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*entity.Lesson
	for rows.Next() {
		lesson := &entity.Lesson{}
		err := rows.Scan(
			&lesson.ID,
			&lesson.ModuleID,
			&lesson.Title,
			&lesson.Description,
			&lesson.Order,
			&lesson.Duration,
			&lesson.CreatedAt,
			&lesson.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, rows.Err()
}

func (r *LessonRepository) Search(ctx context.Context, query string, limit int) ([]*entity.Lesson, error) {
	pattern := "%" + escapeLike(query) + "%"
	sql := `
		SELECT l.id, l.module_id, l.title, COALESCE(l.description, ''), l."order", l.duration, l.created_at, l.updated_at
		FROM lessons l
		JOIN modules m ON m.id = l.module_id
		JOIN courses c ON c.id = m.course_id AND c.is_published
		WHERE l.title ILIKE $1 ESCAPE '\' OR l.description ILIKE $1 ESCAPE '\'
		ORDER BY l.title
		LIMIT $2
	`
	rows, err := r.db.Pool.Query(ctx, sql, pattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*entity.Lesson
	for rows.Next() {
		lesson := &entity.Lesson{}
		err := rows.Scan(
			&lesson.ID,
			&lesson.ModuleID,
			&lesson.Title,
			&lesson.Description,
			&lesson.Order,
			&lesson.Duration,
			&lesson.CreatedAt,
			&lesson.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, rows.Err()
}

func (r *LessonRepository) scanLesson(row pgx.Row) (*entity.Lesson, error) {
	lesson := &entity.Lesson{}
	err := row.Scan(
		&lesson.ID,
		&lesson.ModuleID,
		&lesson.Title,
		&lesson.Description,
		&lesson.Order,
		&lesson.Duration,
		&lesson.CreatedAt,
		&lesson.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return lesson, nil
}
