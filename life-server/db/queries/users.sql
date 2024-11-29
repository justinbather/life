-- name: CreateUser :one
INSERT INTO users (id, username, password) VALUES ($1, $2, $3) RETURNING id, username, password;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUsernameAndPass :one
SELECT * FROM users WHERE username = $1 AND password = $2;
