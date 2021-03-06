// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: comment.sql

package db

import (
	"context"
	"time"
)

const createComment = `-- name: CreateComment :execlastid
insert into comment (user_id, video_id, content)
values (?, ?, ?)
`

type CreateCommentParams struct {
	UserID  int64  `json:"user_id"`
	VideoID int64  `json:"video_id"`
	Content string `json:"content"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createComment, arg.UserID, arg.VideoID, arg.Content)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const deleteComment = `-- name: DeleteComment :exec
delete
from comment
where id = ?
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :one
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
limit 1
`

type GetCommentRow struct {
	ID            int64     `json:"id"`
	Content       string    `json:"content"`
	CreateDate    time.Time `json:"create_date"`
	UserID        int64     `json:"user_id"`
	Name          string    `json:"name"`
	FollowCount   int64     `json:"follow_count"`
	FollowerCount int64     `json:"follower_count"`
}

func (q *Queries) GetComment(ctx context.Context, id int64) (GetCommentRow, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i GetCommentRow
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreateDate,
		&i.UserID,
		&i.Name,
		&i.FollowCount,
		&i.FollowerCount,
	)
	return i, err
}

const getCommentsByVideoId = `-- name: GetCommentsByVideoId :many
select comment.id, comment.content, comment.create_date, user.id as user_id, name, follow_count, follower_count
from comment,
     user
where comment.video_id = ?
  and comment.user_id = user.id
order by comment.create_date desc
`

type GetCommentsByVideoIdRow struct {
	ID            int64     `json:"id"`
	Content       string    `json:"content"`
	CreateDate    time.Time `json:"create_date"`
	UserID        int64     `json:"user_id"`
	Name          string    `json:"name"`
	FollowCount   int64     `json:"follow_count"`
	FollowerCount int64     `json:"follower_count"`
}

func (q *Queries) GetCommentsByVideoId(ctx context.Context, videoID int64) ([]GetCommentsByVideoIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsByVideoId, videoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCommentsByVideoIdRow{}
	for rows.Next() {
		var i GetCommentsByVideoIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.CreateDate,
			&i.UserID,
			&i.Name,
			&i.FollowCount,
			&i.FollowerCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
