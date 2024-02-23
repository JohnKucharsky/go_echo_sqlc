-- name: CreateOrder :one
insert into orders (updated_at, product_id, user_id)
values ($1, $2, $3) returning *;

-- name: GetOrders :many
SELECT orders.id as order_id,
       products.name as product_name,
       products.serial as product_serial,
       users.first_name as user_name,
       users.last_name as user_last_name,
       orders.updated_at as updated_at
FROM orders
         JOIN products ON orders.product_id = products.id
         JOIN users ON orders.user_id = users.id;

-- name: UpdateOrder :one
update orders set product_id=$1, user_id=$2, updated_at=$3
where id = $4 returning *;

-- name: DeleteOrder :exec
delete from orders where id = $1;