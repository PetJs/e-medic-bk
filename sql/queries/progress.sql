-- name: GetProgressByUserAndLesson :one
SELECT * FROM progress WHERE user_id = $1 AND lesson_id = $2;

-- name: UpsertProgress :one
INSERT INTO progress (user_id, lesson_id, is_completed, progress_pct, last_position, completed_at)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (user_id, lesson_id) DO UPDATE SET
    is_completed = EXCLUDED.is_completed,
    progress_pct = EXCLUDED.progress_pct,
    last_position = EXCLUDED.last_position,
    completed_at = EXCLUDED.completed_at,
    updated_at = NOW()
RETURNING *;

-- name: ListProgressByUser :many
SELECT * FROM progress WHERE user_id = $1;
