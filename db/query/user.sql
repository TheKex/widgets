-- name: CreateUser :one
insert into "user"
(
  "username",
  "hashed_password",
  "full_name",
  "email"
)
values
(
  $1,
  $2,
  $3,
  $4
)
returning *;

-- name: GetUser :one
select *
  from "user"
 where id = $1 limit 1;

-- name: GetUserByUsername :one
select *
  from "user"
 where username = $1 limit 1;

-- name: ListUsers :many
select *
  from "user"
 order by id
 limit $1
offset $2;

-- name: UpdateUserPassword :one
update "user" set
  "hashed_password" = $2,
  "password_changed_at" = CURRENT_TIMESTAMP
 where id = $1
returning *;

-- name: DeleteUser :exec
delete from "user"
 where id = $1;


