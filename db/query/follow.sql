-- name: CreateFollow :one
INSERT INTO follow (following_id, follower_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetFollow :one
SELECT *
FROM follow
WHERE following_id = $1
  AND follower_id = $2
LIMIT 1;