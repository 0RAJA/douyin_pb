-- name: CreateUserFollower :exec
insert into user_followers (from_id, to_id)
values (?, ?);

-- name: GetUserFollower :one
select *
from user_followers
where from_id = ?
  and to_id = ?
limit 1;

-- name: DeleteUserFollower :exec
delete
from user_followers
where from_id = ?
  and to_id = ?;

-- name: GetUserFollowersByFromUserID :many
select *
from user_followers
where from_id = ?;

-- name: GetUserFollowersByToUserID :many
select *
from user_followers
where to_id = ?;

-- name: IsExistRelations :one
select count(*) from user_followers
where from_id = ?
and to_id = ?;






