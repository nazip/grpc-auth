CREATE TABLE users (
    id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    email  text,
    password text NOT NULL
);

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    name, email, password
) VALUES (
             $1, $2, $3
         )
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;