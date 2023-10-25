-- +goose Up

CREATE TABLE users (
                       id   BIGSERIAL PRIMARY KEY,
                       name text      NOT NULL,
                       email  text,
                       password text NOT NULL,
                       role int NOT NULL,
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP
);


-- +goose Down

drop table users;

