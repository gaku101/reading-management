-- name: CreateEntry :one
INSERT INTO entries (user_id, amount)
VALUES ($1, $2)
RETURNING *;
-- name: GetEntry :one
SELECT *
FROM entries
WHERE id = $1
LIMIT 1;
-- name: ListEntries :many
SELECT *
FROM entries
WHERE user_id = $1
  AND amount > 0;