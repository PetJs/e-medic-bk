// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type QuizQuestionRepository struct{ db *DB }

func NewQuizQuestionRepository(db *DB) repository.QuizQuestionRepository {
	return &QuizQuestionRepository{db: db}
}

const quizQuestionColumns = `id, lesson_id, type, prompt, "order", created_at, updated_at`
const quizOptionColumns = `id, question_id, text, is_correct, "order"`

// CreateWithOptions inserts the question and its options (if any) inside a
// single transaction, so a multiple_choice question is never left without
// its options if something fails partway through.
func (r *QuizQuestionRepository) CreateWithOptions(ctx context.Context, q *entity.QuizQuestion, options []*entity.QuizOption) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`INSERT INTO lesson_quiz_questions (id, lesson_id, type, prompt, "order", created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		q.ID, q.LessonID, string(q.Type), q.Prompt, q.Order, q.CreatedAt, q.UpdatedAt,
	)
	if err != nil {
		return err
	}

	for _, o := range options {
		_, err = tx.Exec(ctx,
			`INSERT INTO lesson_quiz_options (id, question_id, text, is_correct, "order") VALUES ($1, $2, $3, $4, $5)`,
			o.ID, o.QuestionID, o.Text, o.IsCorrect, o.Order,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *QuizQuestionRepository) GetByID(ctx context.Context, id string) (*entity.QuizQuestion, error) {
	row := r.db.Pool.QueryRow(ctx, `SELECT `+quizQuestionColumns+` FROM lesson_quiz_questions WHERE id = $1`, id)
	q, err := r.scanQuestion(row)
	if err != nil || q == nil {
		return q, err
	}
	options, err := r.listOptions(ctx, q.ID)
	if err != nil {
		return nil, err
	}
	q.Options = options
	return q, nil
}

func (r *QuizQuestionRepository) Update(ctx context.Context, q *entity.QuizQuestion) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE lesson_quiz_questions SET prompt = $2, "order" = $3, updated_at = $4 WHERE id = $1`,
		q.ID, q.Prompt, q.Order, q.UpdatedAt,
	)
	return err
}

func (r *QuizQuestionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM lesson_quiz_questions WHERE id = $1`, id)
	return err
}

func (r *QuizQuestionRepository) ListByLesson(ctx context.Context, lessonID string) ([]*entity.QuizQuestion, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT `+quizQuestionColumns+` FROM lesson_quiz_questions WHERE lesson_id = $1 ORDER BY "order", created_at`,
		lessonID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	questions, err := r.scanQuestions(rows)
	if err != nil {
		return nil, err
	}
	for _, q := range questions {
		options, err := r.listOptions(ctx, q.ID)
		if err != nil {
			return nil, err
		}
		q.Options = options
	}
	return questions, nil
}

func (r *QuizQuestionRepository) CountByLesson(ctx context.Context, lessonID string) (int, error) {
	var count int
	err := r.db.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM lesson_quiz_questions WHERE lesson_id = $1`, lessonID).Scan(&count)
	return count, err
}

func (r *QuizQuestionRepository) listOptions(ctx context.Context, questionID string) ([]*entity.QuizOption, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT `+quizOptionColumns+` FROM lesson_quiz_options WHERE question_id = $1 ORDER BY "order"`,
		questionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []*entity.QuizOption
	for rows.Next() {
		o := &entity.QuizOption{}
		if err := rows.Scan(&o.ID, &o.QuestionID, &o.Text, &o.IsCorrect, &o.Order); err != nil {
			return nil, err
		}
		options = append(options, o)
	}
	return options, rows.Err()
}

func (r *QuizQuestionRepository) scanQuestion(row pgx.Row) (*entity.QuizQuestion, error) {
	q := &entity.QuizQuestion{}
	var qType string
	err := row.Scan(&q.ID, &q.LessonID, &qType, &q.Prompt, &q.Order, &q.CreatedAt, &q.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	q.Type = entity.QuestionType(qType)
	return q, nil
}

func (r *QuizQuestionRepository) scanQuestions(rows pgx.Rows) ([]*entity.QuizQuestion, error) {
	var questions []*entity.QuizQuestion
	for rows.Next() {
		q := &entity.QuizQuestion{}
		var qType string
		if err := rows.Scan(&q.ID, &q.LessonID, &qType, &q.Prompt, &q.Order, &q.CreatedAt, &q.UpdatedAt); err != nil {
			return nil, err
		}
		q.Type = entity.QuestionType(qType)
		questions = append(questions, q)
	}
	return questions, rows.Err()
}
