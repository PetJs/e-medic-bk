// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type ProgressRepository struct{ db *DB }

func NewProgressRepository(db *DB) repository.ProgressRepository { return &ProgressRepository{db: db} }

const progressColumns = `p.id, p.user_id, p.lesson_id, p.is_completed, p.progress_pct, p.last_position, p.completed_at, p.created_at, p.updated_at`

func (r *ProgressRepository) Upsert(ctx context.Context, progress *entity.Progress) error {
	query := `
		INSERT INTO progress (id, user_id, lesson_id, is_completed, progress_pct, last_position, completed_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (user_id, lesson_id) DO UPDATE SET
			is_completed = EXCLUDED.is_completed,
			progress_pct = EXCLUDED.progress_pct,
			last_position = EXCLUDED.last_position,
			completed_at = EXCLUDED.completed_at,
			updated_at = NOW()
	`
	_, err := r.db.Pool.Exec(ctx, query,
		progress.ID,
		progress.UserID,
		progress.LessonID,
		progress.IsCompleted,
		progress.ProgressPct,
		progress.LastPosition,
		progress.CompletedAt,
		progress.CreatedAt,
		progress.UpdatedAt,
	)
	return err
}

func (r *ProgressRepository) GetByUserAndLesson(ctx context.Context, userID, lessonID string) (*entity.Progress, error) {
	query := `SELECT ` + progressColumns + ` FROM progress p WHERE p.user_id = $1 AND p.lesson_id = $2`
	row := r.db.Pool.QueryRow(ctx, query, userID, lessonID)
	return r.scanProgress(row)
}

func (r *ProgressRepository) ListByUser(ctx context.Context, userID string) ([]*entity.Progress, error) {
	query := `
		SELECT ` + progressColumns + `, l.module_id
		FROM progress p
		JOIN lessons l ON l.id = p.lesson_id
		WHERE p.user_id = $1
		ORDER BY p.updated_at DESC
	`
	rows, err := r.db.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanProgressWithModule(rows)
}

func (r *ProgressRepository) ListByUserAndCourse(ctx context.Context, userID, courseID string) ([]*entity.Progress, error) {
	query := `
		SELECT ` + progressColumns + `, l.module_id
		FROM progress p
		JOIN lessons l ON l.id = p.lesson_id
		JOIN modules m ON m.id = l.module_id
		WHERE p.user_id = $1 AND m.course_id = $2
		ORDER BY p.updated_at DESC
	`
	rows, err := r.db.Pool.Query(ctx, query, userID, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanProgressWithModule(rows)
}

func (r *ProgressRepository) GetCourseCompletionStats(ctx context.Context, userID, courseID string) (completed, total int, err error) {
	query := `
		SELECT COUNT(l.id), COUNT(p.id) FILTER (WHERE p.is_completed)
		FROM lessons l
		JOIN modules m ON m.id = l.module_id
		LEFT JOIN progress p ON p.lesson_id = l.id AND p.user_id = $1
		WHERE m.course_id = $2
	`
	err = r.db.Pool.QueryRow(ctx, query, userID, courseID).Scan(&total, &completed)
	return completed, total, err
}

func (r *ProgressRepository) scanProgress(row pgx.Row) (*entity.Progress, error) {
	p := &entity.Progress{}
	err := row.Scan(
		&p.ID,
		&p.UserID,
		&p.LessonID,
		&p.IsCompleted,
		&p.ProgressPct,
		&p.LastPosition,
		&p.CompletedAt,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func (r *ProgressRepository) scanProgressWithModule(rows pgx.Rows) ([]*entity.Progress, error) {
	var list []*entity.Progress
	for rows.Next() {
		p := &entity.Progress{}
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.LessonID,
			&p.IsCompleted,
			&p.ProgressPct,
			&p.LastPosition,
			&p.CompletedAt,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.ModuleID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}
