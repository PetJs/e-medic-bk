// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type QuizAnswerRepository struct{ db *DB }

func NewQuizAnswerRepository(db *DB) repository.QuizAnswerRepository {
	return &QuizAnswerRepository{db: db}
}

const quizAnswerColumns = `id, question_id, user_id, selected_option_id, free_text_body, is_correct, created_at`

func (r *QuizAnswerRepository) Create(ctx context.Context, a *entity.QuizAnswer) error {
	_, err := r.db.Pool.Exec(ctx,
		`INSERT INTO lesson_quiz_answers (id, question_id, user_id, selected_option_id, free_text_body, is_correct, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		a.ID, a.QuestionID, a.UserID, a.SelectedOptionID, a.FreeTextBody, a.IsCorrect, a.CreatedAt,
	)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return repository.ErrAlreadyAnswered
	}
	return err
}

func (r *QuizAnswerRepository) ListByUserAndLesson(ctx context.Context, userID, lessonID string) ([]*entity.QuizAnswer, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT a.id, a.question_id, a.user_id, a.selected_option_id, a.free_text_body, a.is_correct, a.created_at
		 FROM lesson_quiz_answers a
		 JOIN lesson_quiz_questions q ON q.id = a.question_id
		 WHERE a.user_id = $1 AND q.lesson_id = $2`,
		userID, lessonID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanAnswers(rows)
}

func (r *QuizAnswerRepository) ListByQuestion(ctx context.Context, questionID string) ([]*entity.QuizAnswer, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT a.id, a.question_id, a.user_id, a.selected_option_id, a.free_text_body, a.is_correct, a.created_at,
		        u.id, u.email, u.first_name, u.last_name, u.role, u.created_at
		 FROM lesson_quiz_answers a
		 JOIN users u ON u.id = a.user_id
		 WHERE a.question_id = $1
		 ORDER BY a.created_at`,
		questionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var answers []*entity.QuizAnswer
	for rows.Next() {
		a := &entity.QuizAnswer{Author: &entity.User{}}
		err := rows.Scan(
			&a.ID, &a.QuestionID, &a.UserID, &a.SelectedOptionID, &a.FreeTextBody, &a.IsCorrect, &a.CreatedAt,
			&a.Author.ID, &a.Author.Email, &a.Author.FirstName, &a.Author.LastName, &a.Author.Role, &a.Author.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		answers = append(answers, a)
	}
	return answers, rows.Err()
}

func (r *QuizAnswerRepository) CountAnsweredByUser(ctx context.Context, lessonID, userID string) (int, error) {
	var count int
	err := r.db.Pool.QueryRow(ctx,
		`SELECT COUNT(*)
		 FROM lesson_quiz_answers a
		 JOIN lesson_quiz_questions q ON q.id = a.question_id
		 WHERE q.lesson_id = $1 AND a.user_id = $2`,
		lessonID, userID,
	).Scan(&count)
	return count, err
}

func (r *QuizAnswerRepository) scanAnswers(rows pgx.Rows) ([]*entity.QuizAnswer, error) {
	var answers []*entity.QuizAnswer
	for rows.Next() {
		a := &entity.QuizAnswer{}
		if err := rows.Scan(&a.ID, &a.QuestionID, &a.UserID, &a.SelectedOptionID, &a.FreeTextBody, &a.IsCorrect, &a.CreatedAt); err != nil {
			return nil, err
		}
		answers = append(answers, a)
	}
	return answers, rows.Err()
}
