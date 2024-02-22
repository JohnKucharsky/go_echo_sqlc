-- +goose Up
CREATE TABLE products (
   id serial PRIMARY KEY,
   name TEXT NOT NULL,
   serial TEXT,
   updated_at TIMESTAMP NOT NULL
);

-- +goose Down
drop table products;