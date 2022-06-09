-- name: CreateUser :execlastid
INSERT INTO user (name, password)
VALUES (?, ?);

-- name: GetUserByUsername :one
select *
from user
where name = ?
limit 1;

-- name: GetUserByUserID :one
select *
from user
where id = ?
limit 1;

-- name: UpdateUserFollowCount :exec
update user
set follow_count = follow_count + ?
where id = ?;

-- name: UpdateUserFollowerCount :exec
update user
set follower_count = follower_count + ?
where id = ?;

-- name: UserFollowCount :one
select follow_count
from user;

-- name: UserFollowerCount :one
select follower_count
from user;

-- name: AddFavorite :exec
update user
set follower_count = follower_count + 1
where id = ?;

-- name: DeleteFavorite :exec
update user
set follower_count = follower_count - 1
where id = ?

