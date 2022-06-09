// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user_followers.sql

package db

import (
	"context"
)

const createUserFollower = `-- name: CreateUserFollower :exec
insert into user_followers (from_id, to_id)
values (?, ?)
`

type CreateUserFollowerParams struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
}

func (q *Queries) CreateUserFollower(ctx context.Context, arg CreateUserFollowerParams) error {
	_, err := q.db.ExecContext(ctx, createUserFollower, arg.FromID, arg.ToID)
	return err
}

const deleteUserFollower = `-- name: DeleteUserFollower :exec
delete
from user_followers
where from_id = ?
  and to_id = ?
`

type DeleteUserFollowerParams struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
}

func (q *Queries) DeleteUserFollower(ctx context.Context, arg DeleteUserFollowerParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserFollower, arg.FromID, arg.ToID)
	return err
}

const getUserFollower = `-- name: GetUserFollower :one
select id, from_id, to_id
from user_followers
where from_id = ?
  and to_id = ?
limit 1
`

type GetUserFollowerParams struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
}

func (q *Queries) GetUserFollower(ctx context.Context, arg GetUserFollowerParams) (UserFollower, error) {
	row := q.db.QueryRowContext(ctx, getUserFollower, arg.FromID, arg.ToID)
	var i UserFollower
	err := row.Scan(&i.ID, &i.FromID, &i.ToID)
	return i, err
}

const getUserFollowersByFromUserID = `-- name: GetUserFollowersByFromUserID :many
select id, from_id, to_id
from user_followers
where from_id = ?
`

func (q *Queries) GetUserFollowersByFromUserID(ctx context.Context, fromID int64) ([]UserFollower, error) {
	rows, err := q.db.QueryContext(ctx, getUserFollowersByFromUserID, fromID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserFollower{}
	for rows.Next() {
		var i UserFollower
		if err := rows.Scan(&i.ID, &i.FromID, &i.ToID); err != nil {
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

const getUserFollowersByToUserID = `-- name: GetUserFollowersByToUserID :many
select id, from_id, to_id
from user_followers
where to_id = ?
`

func (q *Queries) GetUserFollowersByToUserID(ctx context.Context, toID int64) ([]UserFollower, error) {
	rows, err := q.db.QueryContext(ctx, getUserFollowersByToUserID, toID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserFollower{}
	for rows.Next() {
		var i UserFollower
		if err := rows.Scan(&i.ID, &i.FromID, &i.ToID); err != nil {
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

const isExistRelations = `-- name: IsExistRelations :one
select count(*) from user_followers
where from_id = ?
and to_id = ?
`

type IsExistRelationsParams struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
}

func (q *Queries) IsExistRelations(ctx context.Context, arg IsExistRelationsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, isExistRelations, arg.FromID, arg.ToID)
	var count int64
	err := row.Scan(&count)
	return count, err
}
