-- name: CreateTransfer :one
INSERT INTO transfers (from_user_id, to_user_id, amount)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetTransfer :one
SELECT *
FROM transfers
WHERE id = $1
LIMIT 1;
-- name: ListTransfers :many
SELECT transfers.id,
  from_user_id,
  to_user_id,
  amount,
  transfers.created_at,
  users.username
FROM transfers
  JOIN users ON from_user_id = users.id
  AND to_user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;