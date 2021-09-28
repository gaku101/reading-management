-- name: CreateComment :one
INSERT INTO comments (author, post_id, body)
VALUES ($1, $2, $3)
RETURNING *;