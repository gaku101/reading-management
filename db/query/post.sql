-- name: CreatePost :one
INSERT INTO posts (author, title, body)
VALUES ($1, $2, $3)
RETURNING *;