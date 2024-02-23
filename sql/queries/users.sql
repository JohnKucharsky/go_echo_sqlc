-- name: CreateUser :one
insert into users (first_name, last_name, updated_at)
values ($1, $2, $3) returning *;

-- name: GetUsers :many
select * from users
WHERE CASE WHEN LENGTH(@f_name::text) != 0
    THEN first_name like @f_name::text ELSE TRUE END
  AND
    CASE WHEN LENGTH(@l_name::text) != 0
      THEN last_name like @l_name::text ELSE TRUE END
ORDER BY
    CASE WHEN @sort_order::text = 'asc' THEN
             CASE @order_by::text
                 WHEN 'first_name' THEN first_name
                 WHEN 'last_name' THEN last_name
                 ELSE NULL
                 END
         ELSE
             NULL
        END
        ASC,

    CASE WHEN @sort_order::text = 'desc' THEN
             CASE @order_by::text
                 WHEN 'first_name' THEN first_name
                 WHEN 'last_name' THEN last_name
                 ELSE NULL
                 END
         ELSE
             NULL
        END
        DESC
limit sqlc.narg('limit') offset sqlc.narg('offset')
;

-- name: GetOneUser :one
select * from users where id = $1;

-- name: UpdateUser :one
update users set first_name=$1, last_name=$2, updated_at=$3
where id = $4 returning *;

-- name: DeleteUser :exec
delete from users where id = $1;