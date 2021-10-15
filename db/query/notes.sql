-- name: CreateNote :one
INSERT INTO notes (author, post_id, body, page, line)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: ListNotes :many
SELECT *
FROM notes
WHERE post_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;