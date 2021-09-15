-- name: CreatePostCategory :one
INSERT INTO post_category (post_id, category_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetPostCategory :one
SELECT category.id,
  name
FROM category
  JOIN post_category ON category.id = category_id
  AND post_id = $1
LIMIT 1;