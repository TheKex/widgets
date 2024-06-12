-- name: CreateCounter :one
insert into "counter"
(
  "widget_id",
  "order",
  "value",
  "increment"
)
values
(
  $1,
  $2,
  $3,
  $4
)
returning *;

-- name: GetCounter :one
select *
  from "counter"
 where id = $1;

-- name: ListCounters :many
select *
  from "counter"
 order by id
 limit $1
offset $2;

-- name: UpdateCounter :one
update "counter" set
  "order" = $2,
  "value" = $3,
  "increment" = $4
where id = $1
returning *;

-- name: DeleteCounter :exec
delete from "counter"
 where id = $1;

-- name: UpdateCounterValue :one
update "counter" set
  value = sqlc.arg(value)
where id = sqlc.arg(id)
returning *;
