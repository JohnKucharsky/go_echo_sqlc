-- name: CreateUser :one
insert into users (id, first_name, last_name, updated_at)
values ($1, $2, $3, $4) returning *;
