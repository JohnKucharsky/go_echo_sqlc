-- name: CreateUser :one
insert into users (first_name, last_name, updated_at)
values ($1, $2, $3) returning *;

-- name: GetUsers :many
select * from users;

-- name: GetOneUser :one
select * from users where id = $1;

-- name: UpdateUser :one
update users set first_name=$1, last_name=$2, updated_at=$3
where id = $4 returning *;

-- name: DeleteUser :exec
delete from users where id = $1;