-- name: ListQuestionsByLesson :many
SELECT * FROM questions WHERE lesson_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: GetQuestionByID :one
SELECT * FROM questions WHERE id = $1;

-- name: CreateQuestion :one
INSERT INTO questions (lesson_id, user_id, title, body)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateQuestion :one
UPDATE questions SET title = $2, body = $3, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteQuestion :exec
DELETE FROM questions WHERE id = $1;

-- name: IncrementAnswerCount :exec
UPDATE questions SET answer_count = answer_count + 1 WHERE id = $1;
