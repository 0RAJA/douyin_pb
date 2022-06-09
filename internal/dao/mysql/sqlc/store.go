package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TXer
}

type TXer interface {
	CreateUserFollowerWithTX(ctx context.Context, arg CreateUserFollowerParams) error
	DeleteUserFollowerWithTx(ctx context.Context, arg DeleteUserFollowerParams) error
	CreateCommentWithTx(ctx context.Context, arg CreateCommentParams) (GetCommentRow, error)
	DeleteCommentWithTx(ctx context.Context, commentID, videoID int64) error
	CreateUserVideoWithTx(ctx context.Context, arg CreateUserVideoParams) error
	DeleteUserVideoWithTx(ctx context.Context, arg DeleteUserVideoParams) error
}

type SqlStore struct {
	*Queries
	*sql.DB
}

// 通过事务执行回调函数
func (store *SqlStore) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := store.WithTx(tx) // 使用开启的事务创建一个查询
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); err != nil {
			return fmt.Errorf("tx err:%v,rb err:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// CreateUserFollowerWithTX 通过事务创建关注记录并增加关注和被关注数量
func (store *SqlStore) CreateUserFollowerWithTX(ctx context.Context, arg CreateUserFollowerParams) error {
	return store.execTx(ctx, func(queries *Queries) error {
		var err error
		doThat := func(f func()) {
			if err != nil {
				return
			}
			f()
		}
		doThat(func() {
			err = queries.CreateUserFollower(ctx, arg)
		})
		doThat(func() {
			err = queries.UpdateUserFollowCount(ctx, UpdateUserFollowCountParams{
				FollowCount: 1,
				ID:          arg.FromID,
			})
		})
		doThat(func() {
			err = queries.UpdateUserFollowerCount(ctx, UpdateUserFollowerCountParams{
				FollowerCount: 1,
				ID:            arg.ToID,
			})
		})
		return err
	})
}

// DeleteUserFollowerWithTx 通过事务删除关注关系并减少关注和被关注数量
func (store *SqlStore) DeleteUserFollowerWithTx(ctx context.Context, arg DeleteUserFollowerParams) error {
	return store.execTx(ctx, func(queries *Queries) error {
		var err error
		doThat := func(f func()) {
			if err != nil {
				return
			}
			f()
		}
		doThat(func() {
			if _, err1 := queries.GetUserFollower(ctx, GetUserFollowerParams(arg)); err1 != nil {
				err = err1
			}
		})
		doThat(func() {
			err = queries.DeleteUserFollower(ctx, arg)
		})
		doThat(func() {
			err = queries.UpdateUserFollowCount(ctx, UpdateUserFollowCountParams{
				FollowCount: -1,
				ID:          arg.FromID,
			})
		})
		doThat(func() {
			err = queries.UpdateUserFollowerCount(ctx, UpdateUserFollowerCountParams{
				FollowerCount: -1,
				ID:            arg.ToID})
		})
		return err
	})
}

// CreateCommentWithTx 通过事务创建对视频的评论以及增加其评论数量，返回操作后评论的信息
func (store *SqlStore) CreateCommentWithTx(ctx context.Context, arg CreateCommentParams) (GetCommentRow, error) {
	var result GetCommentRow
	err := store.execTx(ctx, func(queries *Queries) error {
		commentID, err := queries.CreateComment(ctx, arg)
		if err != nil {
			return err
		}
		if err := queries.UpdateVideoCommentCount(ctx, UpdateVideoCommentCountParams{
			CommentCount: 1,
			ID:           arg.VideoID,
		}); err != nil {
			return err
		}
		result, err = queries.GetComment(ctx, commentID)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

// DeleteCommentWithTx 通过事务删除视频的评论以及减少其评论数量
func (store *SqlStore) DeleteCommentWithTx(ctx context.Context, commentID, videoID int64) error {
	return store.execTx(ctx, func(queries *Queries) error {
		var err error
		doThat := func(f func()) {
			if err != nil {
				return
			}
			f()
		}
		doThat(func() {
			err = queries.DeleteComment(ctx, commentID)
		})
		doThat(func() {
			err = queries.UpdateVideoCommentCount(ctx, UpdateVideoCommentCountParams{
				CommentCount: -1,
				ID:           videoID,
			})
		})
		return err
	})
}

// CreateUserVideoWithTx 通过事务创建用户对视频的点赞，以及增加对应视频的点赞数
func (store *SqlStore) CreateUserVideoWithTx(ctx context.Context, arg CreateUserVideoParams) error {
	return store.execTx(ctx, func(queries *Queries) error {
		var err error
		doThat := func(f func()) {
			if err != nil {
				return
			}
			f()
		}
		doThat(func() {
			_, err = queries.CreateUserVideo(ctx, arg)
		})
		doThat(func() {
			err = queries.UpdateVideoFavoriteCount(ctx, UpdateVideoFavoriteCountParams{
				FavoriteCount: 1,
				ID:            arg.VideoID,
			})
		})
		return err
	})
}

// DeleteUserVideoWithTx 删除用户对视频的点赞信息，同时减少对应点赞数
func (store *SqlStore) DeleteUserVideoWithTx(ctx context.Context, arg DeleteUserVideoParams) error {
	return store.execTx(ctx, func(queries *Queries) error {
		var err error
		doThat := func(f func()) {
			if err != nil {
				return
			}
			f()
		}
		doThat(func() {
			err = queries.DeleteUserVideo(ctx, arg)
		})
		doThat(func() {
			err = queries.UpdateVideoFavoriteCount(ctx, UpdateVideoFavoriteCountParams{
				FavoriteCount: -1,
				ID:            arg.VideoID,
			})
		})
		return err
	})
}
