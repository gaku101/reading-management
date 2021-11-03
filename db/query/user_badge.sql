-- name: CreateUserBadge :one
INSERT INTO user_badge (user_id, badge_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetUserBadge :one
SELECT badge.id,
  name
FROM badge
  JOIN user_badge ON badge.id = badge_id
  AND user_id = $1
LIMIT 1;
-- name: UpdateUserBadge :one
UPDATE user_badge
SET badge_id = $2
WHERE user_id = $1
RETURNING *;
-- name: DeleteUserBadge :exec
DELETE FROM user_badge
WHERE user_id = $1;