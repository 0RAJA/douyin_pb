-- name: CreateUserVideo :execlastid
insert into user_videos (user_id, video_id)
values (?, ?);

-- name: GetUserVideo :one
select *
from user_videos
where user_id = ?
  and video_id = ?;

-- name: DeleteUserVideo :exec
delete
from user_videos
where user_id = ?
  and video_id = ?;

-- name: GetVideoByUserId :many
select *
from user_videos
where user_id = ?;




-- name: GetVideosFavorite :many
select videos.id, videos.play_url, videos.cover_url,videos.favorite_count,videos.comment_count,videos.title, user.id as user_id, name, follow_count, follower_count
from videos,
     user
where videos.id = ?
  and videos.user_id = user.id;

-- name: DeleteFavoriteByUserId :exec
delete
from user_videos
where user_id = ?;