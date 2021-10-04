-- name: CreateComment :one
INSERT INTO comments (author, post_id, body)
VALUES ($1, $2, $3)
RETURNING *;
-- name: ListComments :many
SELECT *
FROM comments
WHERE post_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;
-- name: DeleteComments :one
DELETE FROM comments
WHERE post_id = $1
RETURNING *;
-- name: GetCommentsId :many
SELECT id
FROM comments
WHERE post_id = $1;
-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;
-- name: GetComment :one
SELECT *
FROM comments
WHERE id = $1;