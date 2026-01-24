-- name: ListContentsByLesson :many
SELECT * FROM contents WHERE lesson_id = $1 ORDER BY "order";

-- name: GetContentByID :one
SELECT * FROM contents WHERE id = $1;

-- name: CreateContent :one
INSERT INTO contents (lesson_id, type, title, storage_key, mime_type, size, duration, "order")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: DeleteContent :exec
DELETE FROM contents WHERE id = $1;
