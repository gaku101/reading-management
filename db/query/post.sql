-- name: CreatePost :one
INSERT INTO posts (
    author,
    title,
    book_author,
    book_image,
    book_page,
    book_page_read
  )
VALUES ($1, $2, $3, $4, $5, $6)
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
SET book_page_read = $2
WHERE id = $1
RETURNING *;
-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
-- name: ListMyAllPosts :many
SELECT *
FROM posts
WHERE author = $1;