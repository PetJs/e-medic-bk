-- name: GetCourseByID :one
SELECT * FROM courses WHERE id = $1;

-- name: GetCourseBySlug :one
SELECT * FROM courses WHERE slug = $1;

-- name: CreateCourse :one
INSERT INTO courses (title, description, slug, cover_image, author_id, is_published)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateCourse :one
UPDATE courses SET
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    cover_image = COALESCE($4, cover_image),
    is_published = COALESCE($5, is_published),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses WHERE id = $1;

-- name: ListCourses :many
SELECT * FROM courses
WHERE is_published = true
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListCoursesByAuthor :many
SELECT * FROM courses
WHERE author_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
