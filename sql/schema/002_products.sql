-- +goose Up
CREATE TABLE products (
   id serial PRIMARY KEY,
   name TEXT NOT NULL,
   serial TEXT,
   updated_at timestamptz NOT NULL
);

-- +goose Down
drop table products;