-- +goose Up
CREATE TABLE users (
   id serial PRIMARY KEY,
   first_name TEXT NOT NULL unique,
   last_name TEXT,
   updated_at timestamptz NOT NULL
);

-- +goose Down
drop table users;