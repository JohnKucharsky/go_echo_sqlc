// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: products.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createProduct = `-- name: CreateProduct :one
insert into products (name, serial, updated_at)
values ($1, $2, $3) returning id, name, serial, updated_at
`

type CreateProductParams struct {
	Name      string
	Serial    sql.NullString
	UpdatedAt time.Time
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct, arg.Name, arg.Serial, arg.UpdatedAt)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Serial,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
delete from products where id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getOneProduct = `-- name: GetOneProduct :one
select id, name, serial, updated_at from products where id = $1
`

func (q *Queries) GetOneProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getOneProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Serial,
		&i.UpdatedAt,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
select id, name, serial, updated_at from products
`

func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Serial,
			&i.UpdatedAt,
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

const updateProduct = `-- name: UpdateProduct :one
update products set name=$1, serial=$2, updated_at=$3
where id = $4 returning id, name, serial, updated_at
`

type UpdateProductParams struct {
	Name      string
	Serial    sql.NullString
	UpdatedAt time.Time
	ID        int32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Name,
		arg.Serial,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Serial,
		&i.UpdatedAt,
	)
	return i, err
}
