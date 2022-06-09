package db_test

import (
	"context"
	"testing"

	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/stretchr/testify/require"
)

func TestQueries_CreateUserVideo(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				video := testQueriesCreateVideo(t)
				user := testCreateUser(t)
				id, err := dao.Group.Mysql.CreateUserVideo(context.Background(), db.CreateUserVideoParams{
					UserID:  user.ID,
					VideoID: video.ID,
				})
				require.NoError(t, err)
				require.NotEmpty(t, id)
				result, err := dao.Group.Mysql.GetUserVideo(context.Background(), db.GetUserVideoParams{
					UserID:  user.ID,
					VideoID: video.ID,
				})
				require.NoError(t, err)
				require.Equal(t, result.ID, id)
				id, err = dao.Group.Mysql.CreateUserVideo(context.Background(), db.CreateUserVideoParams{
					UserID:  user.ID,
					VideoID: video.ID,
				})
				require.Error(t, err)
				require.Zero(t, id)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testQueriesCreateUserVideo(t *testing.T) db.UserVideo {
	video := testQueriesCreateVideo(t)
	user := testCreateUser(t)
	id, err := dao.Group.Mysql.CreateUserVideo(context.Background(), db.CreateUserVideoParams{
		UserID:  user.ID,
		VideoID: video.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, id)
	result, err := dao.Group.Mysql.GetUserVideo(context.Background(), db.GetUserVideoParams{
		UserID:  user.ID,
		VideoID: video.ID,
	})
	require.NoError(t, err)
	require.Equal(t, result.ID, id)
	return result
}

func TestQueries_DeleteUserVideo(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				userVideo := testQueriesCreateUserVideo(t)
				require.NoError(t, dao.Group.Mysql.DeleteUserVideo(context.Background(), db.DeleteUserVideoParams{
					UserID:  userVideo.UserID,
					VideoID: userVideo.VideoID,
				}))
				result, err := dao.Group.Mysql.GetUserVideo(context.Background(), db.GetUserVideoParams{
					UserID:  userVideo.UserID,
					VideoID: userVideo.VideoID,
				})
				require.Error(t, err)
				require.Empty(t, result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestQueries_GetUserVideo(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				video := testQueriesCreateVideo(t)
				user := testCreateUser(t)
				id, err := dao.Group.Mysql.CreateUserVideo(context.Background(), db.CreateUserVideoParams{
					UserID:  user.ID,
					VideoID: video.ID,
				})
				require.NoError(t, err)
				require.NotEmpty(t, id)
				result, err := dao.Group.Mysql.GetUserVideo(context.Background(), db.GetUserVideoParams{
					UserID:  user.ID,
					VideoID: video.ID,
				})
				require.NoError(t, err)
				require.Equal(t, result.UserID, user.ID)
				require.Equal(t, result.VideoID, video.ID)
				require.NoError(t, dao.Group.Mysql.DeleteUserVideo(context.Background(), db.DeleteUserVideoParams{
					UserID:  user.ID,
					VideoID: video.ID,
				}))
				result, err = dao.Group.Mysql.GetUserVideo(context.Background(), db.GetUserVideoParams{
					UserID:  user.ID,
					VideoID: video.ID,
				})
				require.Error(t, err)
				require.Empty(t, result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}
