-- name: CreatePostFavorite :one
INSERT INTO post_favorites (post_id, user_id)
VALUES ($1, $2)
RETURNING *;