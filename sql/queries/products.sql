-- name: CreateProduct :one
insert into products (name, serial, updated_at)
values ($1, $2, $3) returning *;

-- name: GetProducts :many
select * from products;

-- name: GetOneProduct :one
select * from products where id = $1;

-- name: UpdateProduct :one
update products set name=$1, serial=$2, updated_at=$3
where id = $4 returning *;

-- name: DeleteProduct :exec
delete from products where id = $1;