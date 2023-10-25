CREATE TABLE users (
    id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    email  text,
    password text NOT NULL,
    role int NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    name, email, password, role, created_at
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
set name = $2,
    email = $3,
    password = $4,
    role = $5
WHERE id = $1
RETURNING *;