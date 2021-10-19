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
-- name: ListFollow :many
SELECT users.id,
  username,
  profile,
  image
FROM follow
  JOIN users ON following_id = users.id
  AND follower_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
-- name: DeleteFollow :exec
DELETE FROM follow
WHERE following_id = $1
  AND follower_id = $2;
-- name: DeleteFollows :exec
DELETE FROM follow
WHERE following_id = $1
  OR follower_id = $1;