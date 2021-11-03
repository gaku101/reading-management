-- name: CreateBadge :one
INSERT INTO badge (name)
VALUES ($1)
RETURNING *;
-- name: GetBadge :one
SELECT *
FROM badge
WHERE id = $1
LIMIT 1;
-- name: ListBadges :many
SELECT *
FROM badge
ORDER BY id;