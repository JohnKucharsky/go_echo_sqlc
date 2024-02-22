// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: orders.sql

package database

import (
	"context"
	"time"
)

const createOrder = `-- name: CreateOrder :one
insert into orders (updated_at, product_id, user_id)
values ($1, $2, $3) returning id, updated_at, product_id, user_id
`

type CreateOrderParams struct {
	UpdatedAt time.Time
	ProductID int32
	UserID    int32
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.UpdatedAt, arg.ProductID, arg.UserID)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UpdatedAt,
		&i.ProductID,
		&i.UserID,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
delete from orders where id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getOneOrder = `-- name: GetOneOrder :one
select id, updated_at, product_id, user_id from orders where id = $1
`

func (q *Queries) GetOneOrder(ctx context.Context, id int32) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOneOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UpdatedAt,
		&i.ProductID,
		&i.UserID,
	)
	return i, err
}

const getOrders = `-- name: GetOrders :many
select id, updated_at, product_id, user_id from orders
`

func (q *Queries) GetOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UpdatedAt,
			&i.ProductID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :one
update orders set product_id=$1, user_id=$2, updated_at=$3
where id = $4 returning id, updated_at, product_id, user_id
`

type UpdateOrderParams struct {
	ProductID int32
	UserID    int32
	UpdatedAt time.Time
	ID        int32
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder,
		arg.ProductID,
		arg.UserID,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UpdatedAt,
		&i.ProductID,
		&i.UserID,
	)
	return i, err
}