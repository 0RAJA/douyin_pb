package db_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestSqlStore_CreateUserFollowerWithTX(t *testing.T) {
	defaultBackground := context.Background()
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good: one to many",
			f: func() {
				fromUser := testCreateUser(t)
				nums := int(utils.RandomInt(10, 30))
				toUsers := make([]db.User, 0, nums)
				for i := 0; i < nums; i++ {
					toUsers = append(toUsers, testCreateUser(t))
				}
				for i := range toUsers {
					err := dao.Group.Mysql.CreateUserFollowerWithTX(defaultBackground, db.CreateUserFollowerParams{
						FromID: fromUser.ID,
						ToID:   toUsers[i].ID,
					})
					require.NoError(t, err)
				}
				user1, err := testQueriesGetUserByUserID(fromUser.ID)
				require.NoError(t, err)
				require.EqualValues(t, user1.FollowCount, nums)
				for i := range toUsers {
					user2, err := testQueriesGetUserByUserID(toUsers[i].ID)
					require.NoError(t, err)
					require.EqualValues(t, user2.FollowerCount, 1)
				}
			},
		},
		{
			name: "good: many to one",
			f: func() {
				toUser := testCreateUser(t)
				nums := int(utils.RandomInt(10, 30))
				fromUsers := make([]db.User, 0, nums)
				for i := 0; i < nums; i++ {
					fromUsers = append(fromUsers, testCreateUser(t))
				}
				for i := range fromUsers {
					err := dao.Group.Mysql.CreateUserFollowerWithTX(defaultBackground, db.CreateUserFollowerParams{
						FromID: fromUsers[i].ID,
						ToID:   toUser.ID,
					})
					require.NoError(t, err)
				}
				user1, err := testQueriesGetUserByUserID(toUser.ID)
				require.NoError(t, err)
				require.EqualValues(t, user1.FollowerCount, nums)
				for i := range fromUsers {
					user2, err := testQueriesGetUserByUserID(fromUsers[i].ID)
					require.NoError(t, err)
					require.EqualValues(t, user2.FollowCount, 1)
				}
			},
		},
		{
			name: "good: repeatedly opt",
			f: func() {
				fromUser := testCreateUser(t)
				toUser := testCreateUser(t)
				var err error
				err = dao.Group.Mysql.CreateUserFollowerWithTX(context.Background(), db.CreateUserFollowerParams{
					FromID: fromUser.ID,
					ToID:   toUser.ID,
				})
				require.NoError(t, err)
				err = dao.Group.Mysql.CreateUserFollowerWithTX(context.Background(), db.CreateUserFollowerParams{
					FromID: fromUser.ID,
					ToID:   toUser.ID,
				})
				require.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testSqlStoreCreateUserFollowerWithTX(t *testing.T) (fromID, toID int64) {
	user1 := testCreateUser(t)
	user2 := testCreateUser(t)
	err := dao.Group.Mysql.CreateUserFollowerWithTX(context.Background(), db.CreateUserFollowerParams{
		FromID: user1.ID,
		ToID:   user2.ID,
	})
	require.NoError(t, err)
	return user1.ID, user2.ID
}

func TestSqlStore_DeleteUserFollowerWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				fromID, toID := testSqlStoreCreateUserFollowerWithTX(t)
				err := dao.Group.Mysql.DeleteUserFollowerWithTx(context.Background(), db.DeleteUserFollowerParams{
					FromID: fromID,
					ToID:   toID,
				})
				require.NoError(t, err)
				_, err = dao.Group.Mysql.GetUserFollower(context.Background(), db.GetUserFollowerParams{
					FromID: fromID,
					ToID:   toID,
				})
				require.Error(t, err, sql.ErrNoRows)
				fromUser, err := testQueriesGetUserByUserID(fromID)
				require.NoError(t, err)
				require.Zero(t, fromUser.FollowCount)
				toUser, err := testQueriesGetUserByUserID(toID)
				require.NoError(t, err)
				require.Zero(t, toUser.FollowerCount)
			},
		},
		{
			name: "good: delete repeatedly",
			f: func() {
				fromID, toID := testSqlStoreCreateUserFollowerWithTX(t)
				err := dao.Group.Mysql.DeleteUserFollowerWithTx(context.Background(), db.DeleteUserFollowerParams{
					FromID: fromID,
					ToID:   toID,
				})
				require.NoError(t, err)
				err = dao.Group.Mysql.DeleteUserFollowerWithTx(context.Background(), db.DeleteUserFollowerParams{
					FromID: fromID,
					ToID:   toID,
				})
				require.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestSqlStore_CreateCommentWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				user := testCreateUser(t)
				video := testQueriesCreateVideo(t)
				comment, err := dao.Group.Mysql.CreateCommentWithTx(context.Background(), db.CreateCommentParams{
					UserID:  user.ID,
					VideoID: video.ID,
					Content: utils.RandomString(100),
				})
				require.NoError(t, err)
				require.NotEmpty(t, comment)
				require.Equal(t, comment.UserID, user.ID)
				result, err := dao.Group.Mysql.GetVideoByID(context.Background(), video.ID)
				require.NoError(t, err)
				require.NotEmpty(t, result)
				require.EqualValues(t, result.CommentCount, 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestSqlStore_DeleteCommentWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				user := testCreateUser(t)
				video := testQueriesCreateVideo(t)
				comment, err := dao.Group.Mysql.CreateCommentWithTx(context.Background(), db.CreateCommentParams{
					UserID:  user.ID,
					VideoID: video.ID,
					Content: utils.RandomString(100),
				})
				require.NoError(t, err)
				require.NotEmpty(t, comment)
				require.Equal(t, comment.UserID, user.ID)
				result, err := dao.Group.Mysql.GetVideoByID(context.Background(), video.ID)
				require.NoError(t, err)
				require.NotEmpty(t, result)
				require.EqualValues(t, result.CommentCount, 1)
				require.NoError(t, dao.Group.Mysql.DeleteCommentWithTx(context.Background(), comment.ID, video.ID))
				result, err = dao.Group.Mysql.GetVideoByID(context.Background(), video.ID)
				require.NoError(t, err)
				require.NotEmpty(t, result)
				require.EqualValues(t, result.CommentCount, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestSqlStore_CreateUserVideoWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				user := testCreateUser(t)
				nums := int(utils.RandomInt(0, 30))
				resultMap := make(map[int64]db.GetVideoByIDRow, nums)
				for i := 0; i < nums; i++ {
					video := testQueriesCreateVideo(t)
					require.NoError(t, dao.Group.Mysql.CreateUserVideoWithTx(context.Background(), db.CreateUserVideoParams{
						UserID:  user.ID,
						VideoID: video.ID,
					}))
					video.FavoriteCount = 1
					resultMap[video.ID] = video
				}
				results, err := dao.Group.Mysql.GetVideosByUserVideo(context.Background(), user.ID)
				require.NoError(t, err)
				require.NotNil(t, results)
				require.Len(t, results, nums)
				for i := range results {
					video, ok := resultMap[results[i].ID]
					require.True(t, ok)
					require.EqualValues(t, video, results[i])
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestSqlStore_DeleteUserVideoWithTx(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				user := testCreateUser(t)
				nums := int(utils.RandomInt(0, 30))
				resultMap := make(map[int64]db.GetVideoByIDRow, nums)
				for i := 0; i < nums; i++ {
					video := testQueriesCreateVideo(t)
					require.NoError(t, dao.Group.Mysql.CreateUserVideoWithTx(context.Background(), db.CreateUserVideoParams{
						UserID:  user.ID,
						VideoID: video.ID,
					}))
					video.FavoriteCount = 1
					resultMap[video.ID] = video
				}
				results, err := dao.Group.Mysql.GetVideosByUserVideo(context.Background(), user.ID)
				require.NoError(t, err)
				require.NotNil(t, results)
				require.Len(t, results, nums)
				for i := range results {
					video, ok := resultMap[results[i].ID]
					require.True(t, ok)
					require.EqualValues(t, video, results[i])
				}
				for videoID, v := range resultMap {
					require.NoError(t, dao.Group.Mysql.DeleteUserVideoWithTx(context.Background(), db.DeleteUserVideoParams{
						UserID:  v.UserID,
						VideoID: v.ID,
					}))
					v.FavoriteCount = 0
					resultMap[videoID] = v
				}
				results, err = dao.Group.Mysql.GetVideosByUserVideo(context.Background(), user.ID)
				require.NoError(t, err)
				require.NotNil(t, results)
				require.Len(t, results, nums)
				for i := range results {
					video, ok := resultMap[results[i].ID]
					require.True(t, ok)
					require.EqualValues(t, video, results[i])
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}
