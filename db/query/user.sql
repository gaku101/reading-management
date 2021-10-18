-- name: CreateUser :one
INSERT INTO users (username, hashed_password, email, profile, image)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;
-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: UpdateUser :one
UPDATE users
SET profile = $2
WHERE id = $1
RETURNING *;
-- name: UpdateUserImage :one
UPDATE users
SET image = $2
WHERE username = $1
RETURNING *;
-- name: GetUserImage :one
SELECT image
FROM users
WHERE username = $1
LIMIT 1;
-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;