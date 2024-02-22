-- name: CreateOrder :one
insert into orders (updated_at, product_id, user_id)
values ($1, $2, $3) returning *;

-- name: GetOrders :many
select * from orders;

-- name: GetOneOrder :one
select * from orders where id = $1;

-- name: UpdateOrder :one
update orders set product_id=$1, user_id=$2, updated_at=$3
where id = $4 returning *;

-- name: DeleteOrder :exec
delete from orders where id = $1;