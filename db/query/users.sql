-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  hashed_password,
  phone_number,
  is_admin
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1 LIMIT 1;