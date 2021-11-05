-- name: CreateBadge :one
INSERT INTO badge (name)
VALUES ($1)
RETURNING *;