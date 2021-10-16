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
-- name: GetNote :one
SELECT *
FROM notes
WHERE id = $1
LIMIT 1;
-- name: UpdateNote :one
UPDATE notes
SET body = $2,
  page = $3,
  line = $4
WHERE id = $1
RETURNING *;
-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;