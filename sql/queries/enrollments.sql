-- name: GetEnrollmentByUserAndCourse :one
SELECT * FROM enrollments WHERE user_id = $1 AND course_id = $2;

-- name: ListEnrollmentsByUser :many
SELECT * FROM enrollments WHERE user_id = $1 ORDER BY enrolled_at DESC;

-- name: CreateEnrollment :one
INSERT INTO enrollments (user_id, course_id, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteEnrollment :exec
DELETE FROM enrollments WHERE id = $1;
