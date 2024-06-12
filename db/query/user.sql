-- name: CreateUser :one
insert into "user"
(
  "name",
  "email"
)
values
(
  $1,
  $2
)
returning *;

-- name: GetUser :one
select *
  from "user"
 where id = $1 limit 1;

-- name: ListUsers :many
select *
  from "user"
 order by id
 limit $1
offset $2;

-- name: UpdateUser :one
update "user"
   set "name" = $2
 where id = $1
returning *;

-- name: DeleteUser :exec
delete from "user"
 where id = $1;


