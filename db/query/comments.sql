-- name: CreateComment :one
INSERT INTO comments (author, post_id, body)
VALUES ($1, $2, $3)
RETURNING *;
-- name: ListComments :many
SELECT *
FROM comments
WHERE post_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;