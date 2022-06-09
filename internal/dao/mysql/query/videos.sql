-- name: CreateVideo :execlastid
insert into videos (user_id, title, play_url, cover_url)
values (?, ?, ?, ?);

-- name: GetVideoByID :one
select videos.id,
       title,
       play_url,
       cover_url,
       favorite_count,
       comment_count,
       created_at,
       user.id as user_id,
       name,
       follow_count,
       follower_count
from videos,
     user
where videos.id = ?
  and videos.user_id = user.id
limit 1;

-- name: GetVideosByUserID :many
select videos.id,
       title,
       play_url,
       cover_url,
       favorite_count,
       comment_count,
       created_at,
       user.id                                        as user_id,
       name,
       follow_count,
       follower_count,
       exists(select *
              from user_videos
              where user_videos.user_id = ?
                and user_videos.video_id = videos.id) as is_favorite_video,
       exists(select * from user_followers where user_followers.from_id = ? and user_followers.to_id = videos.user_id
           )                                          as is_favorite_user
from videos,
     user
where videos.user_id = ?
  and videos.user_id = user.id
order by videos.created_at desc;

-- name: GetVideosByDate :many
select videos.id,
       title,
       play_url,
       cover_url,
       favorite_count,
       comment_count,
       created_at,
       user.id                                        as user_id,
       name,
       follow_count,
       follower_count,
       exists(select *
              from user_videos
              where user_videos.user_id = ?
                and user_videos.video_id = videos.id) as is_favorite_video,
       exists(select * from user_followers where user_followers.from_id = ? and user_followers.to_id = videos.user_id
           )                                          as is_favorite_user
from videos,
     user
where videos.created_at <= ?
  and videos.user_id = user.id
order by videos.created_at desc
limit ?;

-- name: DeleteVideo :exec
delete
from videos
where id = ?;

-- name: UpdateVideoFavoriteCount :exec
update videos
set favorite_count = favorite_count +/**/ ?
where id = ?;

-- name: UpdateVideoCommentCount :exec
update videos
set comment_count = comment_count + ?
where id = ?;

-- name: GetVideosByUserVideo :many
select videos.id,
       title,
       play_url,
       cover_url,
       favorite_count,
       comment_count,
       created_at,
       user.id as user_id,
       name,
       follow_count,
       follower_count
from videos,
     user,
     user_videos
where user_videos.user_id = ?
  and videos.id = user_videos.video_id
  and user.id = videos.user_id
order by user_videos.id desc;


