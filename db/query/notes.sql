-- name: CreateNote :one
INSERT INTO notes (author, post_id, body, page, line)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;