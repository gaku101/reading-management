-- name: CreateCategory :one
INSERT INTO category (name)
VALUES ($1)
RETURNING *;
-- name: GetCategory :one
SELECT *
FROM category
WHERE id = $1
LIMIT 1;