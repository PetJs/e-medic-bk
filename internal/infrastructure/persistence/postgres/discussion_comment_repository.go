// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type DiscussionCommentRepository struct{ db *DB }

func NewDiscussionCommentRepository(db *DB) repository.DiscussionCommentRepository {
	return &DiscussionCommentRepository{db: db}
}

const discussionCommentSelect = `
	SELECT c.id, c.post_id, c.user_id, c.parent_comment_id, c.body, c.created_at, c.updated_at,
	       u.id, u.email, u.first_name, u.last_name, u.role, u.created_at
	FROM discussion_comments c
	JOIN users u ON u.id = c.user_id
`

func (r *DiscussionCommentRepository) Create(ctx context.Context, comment *entity.DiscussionComment) error {
	query := `
		INSERT INTO discussion_comments (id, post_id, user_id, parent_comment_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		comment.ID,
		comment.PostID,
		comment.UserID,
		comment.ParentCommentID,
		comment.Body,
		comment.CreatedAt,
		comment.UpdatedAt,
	)
	return err
}

func (r *DiscussionCommentRepository) GetByID(ctx context.Context, id string) (*entity.DiscussionComment, error) {
	row := r.db.Pool.QueryRow(ctx, discussionCommentSelect+` WHERE c.id = $1`, id)
	return r.scanComment(row)
}

func (r *DiscussionCommentRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM discussion_comments WHERE id = $1`, id)
	return err
}

func (r *DiscussionCommentRepository) ListByPost(ctx context.Context, postID string) ([]*entity.DiscussionComment, error) {
	rows, err := r.db.Pool.Query(ctx,
		discussionCommentSelect+` WHERE c.post_id = $1 ORDER BY c.created_at ASC`,
		postID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanComments(rows)
}

func (r *DiscussionCommentRepository) scanComment(row pgx.Row) (*entity.DiscussionComment, error) {
	comment := &entity.DiscussionComment{Author: &entity.User{}}
	err := row.Scan(
		&comment.ID, &comment.PostID, &comment.UserID, &comment.ParentCommentID, &comment.Body, &comment.CreatedAt, &comment.UpdatedAt,
		&comment.Author.ID, &comment.Author.Email, &comment.Author.FirstName, &comment.Author.LastName, &comment.Author.Role, &comment.Author.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return comment, nil
}

func (r *DiscussionCommentRepository) scanComments(rows pgx.Rows) ([]*entity.DiscussionComment, error) {
	var comments []*entity.DiscussionComment
	for rows.Next() {
		comment := &entity.DiscussionComment{Author: &entity.User{}}
		err := rows.Scan(
			&comment.ID, &comment.PostID, &comment.UserID, &comment.ParentCommentID, &comment.Body, &comment.CreatedAt, &comment.UpdatedAt,
			&comment.Author.ID, &comment.Author.Email, &comment.Author.FirstName, &comment.Author.LastName, &comment.Author.Role, &comment.Author.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, rows.Err()
}
