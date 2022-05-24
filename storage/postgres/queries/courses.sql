-- name: ListCourses :many
SELECT
    *
FROM
    courses;

-- name: AddCourse :one
INSERT INTO
    courses (id, is_active, course_name, created_at, updated_at)
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;
