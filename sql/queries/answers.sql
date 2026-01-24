-- name: ListAnswersByQuestion :many
SELECT * FROM answers WHERE question_id = $1 ORDER BY created_at;

-- name: GetAnswerByID :one
SELECT * FROM answers WHERE id = $1;

-- name: CreateAnswer :one
INSERT INTO answers (question_id, user_id, body)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAnswer :one
UPDATE answers SET body = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteAnswer :exec
DELETE FROM answers WHERE id = $1;

-- name: MarkBestAnswer :one
UPDATE answers SET is_best = true, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: ClearBestAnswer :exec
UPDATE answers SET is_best = false WHERE question_id = $1;
