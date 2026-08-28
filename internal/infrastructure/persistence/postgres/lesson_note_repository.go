// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type LessonNoteRepository struct{ db *DB }

func NewLessonNoteRepository(db *DB) repository.LessonNoteRepository {
	return &LessonNoteRepository{db: db}
}

const lessonNoteColumns = `id, user_id, lesson_id, body, video_position, created_at, updated_at`

func (r *LessonNoteRepository) Create(ctx context.Context, note *entity.LessonNote) error {
	query := `
		INSERT INTO lesson_notes (id, user_id, lesson_id, body, video_position, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		note.ID, note.UserID, note.LessonID, note.Body, note.VideoPosition, note.CreatedAt, note.UpdatedAt,
	)
	return err
}

func (r *LessonNoteRepository) GetByID(ctx context.Context, id string) (*entity.LessonNote, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+lessonNoteColumns+` FROM lesson_notes WHERE id = $1`, id)
	return r.scanNote(row)
}

func (r *LessonNoteRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM lesson_notes WHERE id = $1`, id)
	return err
}

func (r *LessonNoteRepository) ListByUserAndLesson(ctx context.Context, userID, lessonID string) ([]*entity.LessonNote, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT `+lessonNoteColumns+` FROM lesson_notes WHERE user_id = $1 AND lesson_id = $2 ORDER BY video_position ASC`,
		userID, lessonID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*entity.LessonNote
	for rows.Next() {
		note := &entity.LessonNote{}
		if err := rows.Scan(&note.ID, &note.UserID, &note.LessonID, &note.Body, &note.VideoPosition, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, rows.Err()
}

func (r *LessonNoteRepository) scanNote(row pgx.Row) (*entity.LessonNote, error) {
	note := &entity.LessonNote{}
	err := row.Scan(&note.ID, &note.UserID, &note.LessonID, &note.Body, &note.VideoPosition, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return note, nil
}
