-- name: CreatePostFavorite :one
INSERT INTO post_favorites (post_id, user_id)
VALUES ($1, $2)
RETURNING *;
-- name: ListFavoritePosts :many
SELECT posts.id,
  author,
  title,
  body,
  created_at,
  updated_at
FROM posts
  JOIN post_favorites ON posts.id = post_id
  AND user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;