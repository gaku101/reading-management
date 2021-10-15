-- name: CreatePostFavorite :one
INSERT INTO post_favorites (post_id, user_id)
VALUES ($1, $2)
RETURNING *;
-- name: ListFavoritePosts :many
SELECT posts.id,
  author,
  title,
  book_image,
  created_at,
  updated_at
FROM posts
  JOIN post_favorites ON posts.id = post_id
  AND user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
-- name: GetMyFavoritePost :one
SELECT *
FROM post_favorites
WHERE post_id = $1
  AND user_id = $2
LIMIT 1;
-- name: GetPostFavorite :many
SELECT id
FROM post_favorites
WHERE post_id = $1;
-- name: DeletePostFavorite :one
DELETE FROM post_favorites
WHERE post_id = $1
RETURNING *;