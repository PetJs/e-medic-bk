-- name: ListLessonsByModule :many
SELECT * FROM lessons WHERE module_id = $1 ORDER BY "order";

-- name: GetLessonByID :one
SELECT * FROM lessons WHERE id = $1;

-- name: CreateLesson :one
INSERT INTO lessons (module_id, title, description, "order", duration)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateLesson :one
UPDATE lessons SET
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    "order" = COALESCE($4, "order"),
    duration = COALESCE($5, duration),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteLesson :exec
DELETE FROM lessons WHERE id = $1;
