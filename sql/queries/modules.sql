-- name: ListModulesByCourse :many
SELECT * FROM modules WHERE course_id = $1 ORDER BY "order";

-- name: GetModuleByID :one
SELECT * FROM modules WHERE id = $1;

-- name: CreateModule :one
INSERT INTO modules (course_id, title, description, "order", is_premium)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateModule :one
UPDATE modules SET
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    "order" = COALESCE($4, "order"),
    is_premium = COALESCE($5, is_premium),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteModule :exec
DELETE FROM modules WHERE id = $1;
