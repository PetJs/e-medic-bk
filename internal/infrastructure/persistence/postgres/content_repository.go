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

type ContentRepository struct{ db *DB }

func NewContentRepository(db *DB) repository.ContentRepository { return &ContentRepository{db: db} }

const contentColumns = `id, lesson_id, type, title, storage_key, mime_type, size, duration, "order", created_at, updated_at`

func (r *ContentRepository) Create(ctx context.Context, content *entity.Content) error {
	query := `
		INSERT INTO contents (id, lesson_id, type, title, storage_key, mime_type, size, duration, "order", created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		content.ID,
		content.LessonID,
		string(content.Type),
		content.Title,
		content.StorageKey,
		content.MimeType,
		content.Size,
		content.Duration,
		content.Order,
		content.CreatedAt,
		content.UpdatedAt,
	)
	return err
}

func (r *ContentRepository) GetByID(ctx context.Context, id string) (*entity.Content, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+contentColumns+` FROM contents WHERE id = $1`, id)
	return r.scanContent(row)
}

func (r *ContentRepository) Update(ctx context.Context, content *entity.Content) error {
	query := `
		UPDATE contents
		SET type = $2, title = $3, storage_key = $4, mime_type = $5, size = $6, duration = $7, "order" = $8, updated_at = $9
		WHERE id = $1
	`
	_, err := r.db.Pool.Exec(ctx, query,
		content.ID,
		string(content.Type),
		content.Title,
		content.StorageKey,
		content.MimeType,
		content.Size,
		content.Duration,
		content.Order,
		time.Now(),
	)
	return err
}

func (r *ContentRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM contents WHERE id = $1`, id)
	return err
}

func (r *ContentRepository) ListByLesson(ctx context.Context, lessonID string) ([]*entity.Content, error) {
	query := `SELECT ` + contentColumns + ` FROM contents WHERE lesson_id = $1 ORDER BY "order", created_at`
	rows, err := r.db.Pool.Query(ctx, query, lessonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contents []*entity.Content
	for rows.Next() {
		content := &entity.Content{}
		var contentType string
		err := rows.Scan(
			&content.ID,
			&content.LessonID,
			&contentType,
			&content.Title,
			&content.StorageKey,
			&content.MimeType,
			&content.Size,
			&content.Duration,
			&content.Order,
			&content.CreatedAt,
			&content.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		content.Type = entity.ContentType(contentType)
		contents = append(contents, content)
	}
	return contents, rows.Err()
}

func (r *ContentRepository) scanContent(row pgx.Row) (*entity.Content, error) {
	content := &entity.Content{}
	var contentType string
	err := row.Scan(
		&content.ID,
		&content.LessonID,
		&contentType,
		&content.Title,
		&content.StorageKey,
		&content.MimeType,
		&content.Size,
		&content.Duration,
		&content.Order,
		&content.CreatedAt,
		&content.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	content.Type = entity.ContentType(contentType)
	return content, nil
}
