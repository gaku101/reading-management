-- name: CreateFollow :one
INSERT INTO follow (following_id, follower_id)
VALUES ($1, $2)
RETURNING *;