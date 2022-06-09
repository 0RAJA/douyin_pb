// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
)

const addFavorite = `-- name: AddFavorite :exec
update user
set follower_count = follower_count + 1
where id = ?
`

func (q *Queries) AddFavorite(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, addFavorite, id)
	return err
}

const createUser = `-- name: CreateUser :execlastid
INSERT INTO user (name, password)
VALUES (?, ?)
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createUser, arg.Name, arg.Password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const deleteFavorite = `-- name: DeleteFavorite :exec
update user
set follower_count = follower_count - 1
where id = ?
`

func (q *Queries) DeleteFavorite(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFavorite, id)
	return err
}

const getUserByUserID = `-- name: GetUserByUserID :one
select id, name, password, follow_count, follower_count
from user
where id = ?
limit 1
`

func (q *Queries) GetUserByUserID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUserID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.FollowCount,
		&i.FollowerCount,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
select id, name, password, follow_count, follower_count
from user
where name = ?
limit 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.FollowCount,
		&i.FollowerCount,
	)
	return i, err
}

const updateUserFollowCount = `-- name: UpdateUserFollowCount :exec
update user
set follow_count = follow_count + ?
where id = ?
`

type UpdateUserFollowCountParams struct {
	FollowCount int64 `json:"follow_count"`
	ID          int64 `json:"id"`
}

func (q *Queries) UpdateUserFollowCount(ctx context.Context, arg UpdateUserFollowCountParams) error {
	_, err := q.db.ExecContext(ctx, updateUserFollowCount, arg.FollowCount, arg.ID)
	return err
}

const updateUserFollowerCount = `-- name: UpdateUserFollowerCount :exec
update user
set follower_count = follower_count + ?
where id = ?
`

type UpdateUserFollowerCountParams struct {
	FollowerCount int64 `json:"follower_count"`
	ID            int64 `json:"id"`
}

func (q *Queries) UpdateUserFollowerCount(ctx context.Context, arg UpdateUserFollowerCountParams) error {
	_, err := q.db.ExecContext(ctx, updateUserFollowerCount, arg.FollowerCount, arg.ID)
	return err
}

const userFollowCount = `-- name: UserFollowCount :one
select follow_count
from user
`

func (q *Queries) UserFollowCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, userFollowCount)
	var follow_count int64
	err := row.Scan(&follow_count)
	return follow_count, err
}

const userFollowerCount = `-- name: UserFollowerCount :one
select follower_count
from user
`

func (q *Queries) UserFollowerCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, userFollowerCount)
	var follower_count int64
	err := row.Scan(&follower_count)
	return follower_count, err
}
