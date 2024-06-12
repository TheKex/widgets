-- name: CreateWidget :one
insert into "widget"
(
  "name",
  "user_id"
)
values
(
  $1,
  $2
)
returning *;

-- name: GetWidget :one
select *
  from "widget"
 where id = $1
 limit 1;

-- name: ListWidgets :many
select *
  from "widget"
 order by id
 limit $1
offset $2;

-- name: UpdateWidget :one
update "widget" set
  "name" = $2
 where id = $1
returning *;

-- name: DeleteWidget :exec
delete from "widget"
 where id = $1;

