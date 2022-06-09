package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestQueries_CreateComment(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				movie := testQueriesCreateVideo(t)
				t1 := time.Now()
				commentID, err := dao.Group.Mysql.CreateComment(context.Background(), db.CreateCommentParams{
					UserID:  movie.UserID,
					VideoID: movie.ID,
					Content: utils.RandomString(30),
				})
				require.NoError(t, err)
				require.NotZero(t, commentID)
				result, err := dao.Group.Mysql.GetComment(context.Background(), commentID)
				require.NoError(t, err)
				require.NotEmpty(t, result)
				require.Equal(t, result.UserID, movie.UserID)
				require.True(t, result.CreateDate.Sub(t1) <= time.Second, t1.Sub(result.CreateDate) <= time.Second)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testQueriesCreateComment(t *testing.T) db.GetCommentRow {
	movie := testQueriesCreateVideo(t)
	t1 := time.Now()
	commentID, err := dao.Group.Mysql.CreateComment(context.Background(), db.CreateCommentParams{
		UserID:  movie.UserID,
		VideoID: movie.ID,
		Content: utils.RandomString(30),
	})
	require.NoError(t, err)
	require.NotZero(t, commentID)
	result, err := dao.Group.Mysql.GetComment(context.Background(), commentID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.UserID, movie.UserID)
	require.True(t, result.CreateDate.Sub(t1) <= time.Second, t1.Sub(result.CreateDate) <= time.Second)
	return result
}

func TestQueries_DeleteComment(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				comment := testQueriesCreateComment(t)
				result, err := dao.Group.Mysql.GetComment(context.Background(), comment.ID)
				require.NoError(t, err)
				require.NotEmpty(t, result)
				require.NoError(t, dao.Group.Mysql.DeleteComment(context.Background(), comment.ID))
				result, err = dao.Group.Mysql.GetComment(context.Background(), comment.ID)
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

func TestQueries_GetComment(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good:",
			f: func() {
				comment := testQueriesCreateComment(t)
				result, err := dao.Group.Mysql.GetComment(context.Background(), comment.ID)
				require.NoError(t, err)
				require.Equal(t, result.ID, comment.ID)
				require.Equal(t, result.Name, comment.Name)
				require.Equal(t, result.Content, comment.Content)
				require.NoError(t, dao.Group.Mysql.DeleteComment(context.Background(), comment.ID))
				result, err = dao.Group.Mysql.GetComment(context.Background(), comment.ID)
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

func TestQueries_GetCommentsByVideoId(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				video := testQueriesCreateVideo(t)
				nums := int(utils.RandomInt(1, 100))
				nums2 := int(utils.RandomInt(1, 20))
				resultMap := make(map[int64]bool, nums+nums2)
				for i := 0; i < nums; i++ {
					user := testCreateUser(t)
					commentID, err := dao.Group.Mysql.CreateComment(context.Background(), db.CreateCommentParams{
						UserID:  user.ID,
						VideoID: video.ID,
						Content: utils.RandomString(30),
					})
					require.NoError(t, err)
					resultMap[commentID] = true
				}
				user := testCreateUser(t)
				for i := 0; i < nums2; i++ {
					commentID, err := dao.Group.Mysql.CreateComment(context.Background(), db.CreateCommentParams{
						UserID:  user.ID,
						VideoID: video.ID,
						Content: utils.RandomString(30),
					})
					require.NoError(t, err)
					resultMap[commentID] = true
				}
				results, err := dao.Group.Mysql.GetCommentsByVideoId(context.Background(), video.ID)
				require.NoError(t, err)
				require.NotNil(t, results)
				require.Len(t, results, nums+nums2)
				for i := range results {
					require.True(t, resultMap[results[i].ID])
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
