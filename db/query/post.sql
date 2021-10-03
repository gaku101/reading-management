-- name: CreatePost :one
INSERT INTO posts (author, title, body)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPost :one
SELECT *
FROM posts
WHERE id = $1
LIMIT 1;
-- name: ListMyPosts :many
SELECT *
FROM posts
WHERE author = $1
ORDER BY id
LIMIT $2 OFFSET $3;
-- name: ListPosts :many
SELECT *
FROM posts
WHERE NOT author = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
-- name: UpdatePost :one
UPDATE posts
SET title = $2,
  body = $3
WHERE id = $1
RETURNING *;
-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING *;