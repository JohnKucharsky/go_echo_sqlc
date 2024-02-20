-- +goose Up
CREATE TABLE users (
   id serial PRIMARY KEY,
   first_name TEXT NOT NULL,
   last_name TEXT,
   updated_at TIMESTAMP NOT NULL
);

-- +goose Down
drop table users;