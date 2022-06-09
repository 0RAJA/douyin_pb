-- name: CreateComment :execlastid
insert into comment (user_id, video_id, content)
values (?, ?, ?);

-- name: GetComment :one
select comment.id,
       comment.content,
       comment.create_date,
       user.id as user_id,
       name,
       follow_count,
       follower_count
from comment,
     user
where comment.id = ?
  and comment.user_id = user.id
limit 1;

-- name: DeleteComment :exec
delete
from comment
where id = ?;

-- name: GetCommentsByVideoId :many
select comment.id, comment.content, comment.create_date, user.id as user_id, name, follow_count, follower_count
from comment,
     user
where comment.video_id = ?
  and comment.user_id = user.id
order by comment.create_date desc;
