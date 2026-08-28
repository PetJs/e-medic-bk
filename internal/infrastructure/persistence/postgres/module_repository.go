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

type ModuleRepository struct{ db *DB }

func NewModuleRepository(db *DB) repository.ModuleRepository { return &ModuleRepository{db: db} }

// Lists join lessons to expose per-module lesson counts and total duration.
const moduleSelect = `
	SELECT m.id, m.course_id, m.title, COALESCE(m.description, ''), m."order", m.is_premium,
	       COALESCE(m.cover_image_key, ''), m.cover_image_status,
	       m.created_at, m.updated_at,
	       COUNT(l.id), COALESCE(SUM(l.duration), 0)
	FROM modules m
	LEFT JOIN lessons l ON l.module_id = m.id
`

const moduleGroupBy = ` GROUP BY m.id, m.course_id, m.title, m.description, m."order", m.is_premium, m.cover_image_key, m.cover_image_status, m.created_at, m.updated_at`

func (r *ModuleRepository) Create(ctx context.Context, module *entity.Module) error {
	query := `
		INSERT INTO modules (id, course_id, title, description, "order", is_premium, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		module.ID,
		module.CourseID,
		module.Title,
		module.Description,
		module.Order,
		module.IsPremium,
		module.CreatedAt,
		module.UpdatedAt,
	)
	return err
}

func (r *ModuleRepository) GetByID(ctx context.Context, id string) (*entity.Module, error) {
	row := r.db.Pool.QueryRow(ctx, moduleSelect+` WHERE m.id = $1`+moduleGroupBy, id)
	return r.scanModule(row)
}

func (r *ModuleRepository) Update(ctx context.Context, module *entity.Module) error {
	query := `
		UPDATE modules
		SET title = $2, description = $3, "order" = $4, is_premium = $5, updated_at = $6
		WHERE id = $1
	`
	_, err := r.db.Pool.Exec(ctx, query,
		module.ID,
		module.Title,
		module.Description,
		module.Order,
		module.IsPremium,
		time.Now(),
	)
	return err
}

func (r *ModuleRepository) UpdateCoverImage(ctx context.Context, id, coverImageKey, coverImageStatus string) error {
	query := `UPDATE modules SET cover_image_key = $2, cover_image_status = $3, updated_at = $4 WHERE id = $1`
	_, err := r.db.Pool.Exec(ctx, query, id, coverImageKey, coverImageStatus, time.Now())
	return err
}

func (r *ModuleRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM modules WHERE id = $1`, id)
	return err
}

func (r *ModuleRepository) ListByCourse(ctx context.Context, courseID string) ([]*entity.Module, error) {
	rows, err := r.db.Pool.Query(ctx,
		moduleSelect+` WHERE m.course_id = $1`+moduleGroupBy+` ORDER BY m."order", m.created_at`,
		courseID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanModules(rows)
}

func (r *ModuleRepository) ListAll(ctx context.Context) ([]*entity.Module, error) {
	query := `
		SELECT m.id, m.course_id, m.title, COALESCE(m.description, ''), m."order", m.is_premium,
		       COALESCE(m.cover_image_key, ''), m.cover_image_status,
		       m.created_at, m.updated_at,
		       COUNT(l.id), COALESCE(SUM(l.duration), 0)
		FROM modules m
		JOIN courses c ON c.id = m.course_id AND c.is_published
		LEFT JOIN lessons l ON l.module_id = m.id
	` + moduleGroupBy + ` ORDER BY m."order", m.created_at`
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanModules(rows)
}

func (r *ModuleRepository) Search(ctx context.Context, query string, limit int) ([]*entity.Module, error) {
	pattern := "%" + escapeLike(query) + "%"
	sql := `
		SELECT m.id, m.course_id, m.title, COALESCE(m.description, ''), m."order", m.is_premium,
		       COALESCE(m.cover_image_key, ''), m.cover_image_status,
		       m.created_at, m.updated_at,
		       COUNT(l.id), COALESCE(SUM(l.duration), 0)
		FROM modules m
		JOIN courses c ON c.id = m.course_id AND c.is_published
		LEFT JOIN lessons l ON l.module_id = m.id
		WHERE m.title ILIKE $1 ESCAPE '\' OR m.description ILIKE $1 ESCAPE '\'
	` + moduleGroupBy + ` ORDER BY m.title LIMIT $2`
	rows, err := r.db.Pool.Query(ctx, sql, pattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanModules(rows)
}

func (r *ModuleRepository) scanModule(row pgx.Row) (*entity.Module, error) {
	module := &entity.Module{}
	err := row.Scan(
		&module.ID,
		&module.CourseID,
		&module.Title,
		&module.Description,
		&module.Order,
		&module.IsPremium,
		&module.CoverImageKey,
		&module.CoverImageStatus,
		&module.CreatedAt,
		&module.UpdatedAt,
		&module.LessonCount,
		&module.TotalDuration,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return module, nil
}

func (r *ModuleRepository) scanModules(rows pgx.Rows) ([]*entity.Module, error) {
	var modules []*entity.Module
	for rows.Next() {
		module := &entity.Module{}
		err := rows.Scan(
			&module.ID,
			&module.CourseID,
			&module.Title,
			&module.Description,
			&module.Order,
			&module.IsPremium,
			&module.CoverImageKey,
			&module.CoverImageStatus,
			&module.CreatedAt,
			&module.UpdatedAt,
			&module.LessonCount,
			&module.TotalDuration,
		)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	return modules, rows.Err()
}
