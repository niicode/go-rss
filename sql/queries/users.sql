-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;


-- name: ListUsers :many
SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;


-- name: UpdateUser :one
UPDATE users SET updated_at = $2, name = $3 WHERE id = $1
RETURNING *;

