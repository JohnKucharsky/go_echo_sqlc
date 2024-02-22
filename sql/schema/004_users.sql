-- +goose Up
alter table users add constraint unique_name unique (first_name);

-- +goose Down
alter table users drop constraint unique_name;