-- name: CreateUser :one
INSERT INTO users (username, hashed_password, email, profile, image)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;
-- name: UpdateUser :one
UPDATE users
SET username = $2,
  profile = $3,
  image = $4
WHERE id = $1
RETURNING *;