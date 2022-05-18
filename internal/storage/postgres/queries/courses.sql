-- name: ListCourses :many
SELECT
    *
FROM
    courses;

-- name: AddCourse :one
INSERT INTO
    courses (
        id,
        is_active,
        course_name
    )
VALUES
    ($ 1, $ 2, $ 3) RETURNING *;
