// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type DiscussionPostRepository struct{ db *DB }

func NewDiscussionPostRepository(db *DB) repository.DiscussionPostRepository {
	return &DiscussionPostRepository{db: db}
}

// Joins users for the author and left-joins comments to expose a read-time
// comment count. Pinned posts sort first, then newest first.
const discussionPostSelect = `
	SELECT p.id, p.module_id, p.user_id, p.title, p.body, p.is_pinned, p.created_at, p.updated_at,
	       u.id, u.email, u.first_name, u.last_name, u.role, u.created_at,
	       COUNT(c.id)
	FROM discussion_posts p
	JOIN users u ON u.id = p.user_id
	LEFT JOIN discussion_comments c ON c.post_id = p.id
`

const discussionPostGroupBy = ` GROUP BY p.id, p.module_id, p.user_id, p.title, p.body, p.is_pinned, p.created_at, p.updated_at, u.id, u.email, u.first_name, u.last_name, u.role, u.created_at`

func (r *DiscussionPostRepository) Create(ctx context.Context, post *entity.DiscussionPost) error {
	query := `
		INSERT INTO discussion_posts (id, module_id, user_id, title, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		post.ID,
		post.ModuleID,
		post.UserID,
		post.Title,
		post.Body,
		post.CreatedAt,
		post.UpdatedAt,
	)
	return err
}

func (r *DiscussionPostRepository) GetByID(ctx context.Context, id string) (*entity.DiscussionPost, error) {
	row := r.db.Pool.QueryRow(ctx, discussionPostSelect+` WHERE p.id = $1`+discussionPostGroupBy, id)
	return r.scanPost(row)
}

func (r *DiscussionPostRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM discussion_posts WHERE id = $1`, id)
	return err
}

func (r *DiscussionPostRepository) ListByModule(ctx context.Context, moduleID string, limit, offset int) ([]*entity.DiscussionPost, error) {
	rows, err := r.db.Pool.Query(ctx,
		discussionPostSelect+` WHERE p.module_id = $1`+discussionPostGroupBy+`
		ORDER BY p.is_pinned DESC, p.created_at DESC
		LIMIT $2 OFFSET $3`,
		moduleID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanPosts(rows)
}

func (r *DiscussionPostRepository) CountByModule(ctx context.Context, moduleID string) (int64, error) {
	var count int64
	err := r.db.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM discussion_posts WHERE module_id = $1`, moduleID).Scan(&count)
	return count, err
}

func (r *DiscussionPostRepository) SetPinned(ctx context.Context, id string, pinned bool) error {
	_, err := r.db.Pool.Exec(ctx, `UPDATE discussion_posts SET is_pinned = $2, updated_at = NOW() WHERE id = $1`, id, pinned)
	return err
}

func (r *DiscussionPostRepository) scanPost(row pgx.Row) (*entity.DiscussionPost, error) {
	post := &entity.DiscussionPost{Author: &entity.User{}}
	err := row.Scan(
		&post.ID, &post.ModuleID, &post.UserID, &post.Title, &post.Body, &post.IsPinned, &post.CreatedAt, &post.UpdatedAt,
		&post.Author.ID, &post.Author.Email, &post.Author.FirstName, &post.Author.LastName, &post.Author.Role, &post.Author.CreatedAt,
		&post.CommentCount,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return post, nil
}

func (r *DiscussionPostRepository) scanPosts(rows pgx.Rows) ([]*entity.DiscussionPost, error) {
	var posts []*entity.DiscussionPost
	for rows.Next() {
		post := &entity.DiscussionPost{Author: &entity.User{}}
		err := rows.Scan(
			&post.ID, &post.ModuleID, &post.UserID, &post.Title, &post.Body, &post.IsPinned, &post.CreatedAt, &post.UpdatedAt,
			&post.Author.ID, &post.Author.Email, &post.Author.FirstName, &post.Author.LastName, &post.Author.Role, &post.Author.CreatedAt,
			&post.CommentCount,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, rows.Err()
}
