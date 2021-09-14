-- name: CreatePostCategory :one
INSERT INTO post_category (post_id, category_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetPostCategory :one
SELECT *
FROM post_category
WHERE post_id = $1
LIMIT 1;