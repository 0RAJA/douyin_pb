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

func TestQueries_CreateVideo(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				user := testCreateUser(t)
				params := db.CreateVideoParams{
					UserID:   user.ID,
					Title:    utils.RandomString(10),
					PlayUrl:  utils.RandomString(20),
					CoverUrl: utils.RandomString(20),
				}
				t1 := time.Now()
				videoID, err := dao.Group.Mysql.CreateVideo(context.Background(), params)
				require.NoError(t, err)
				require.NotZero(t, videoID)
				video, err := dao.Group.Mysql.GetVideoByID(context.Background(), videoID)
				require.NoError(t, err)
				require.Equal(t, params.Title, video.Title)
				require.Equal(t, params.UserID, video.UserID)
				require.Equal(t, params.PlayUrl, video.PlayUrl)
				require.True(t, video.CreatedAt.Sub(t1) <= time.Second, t1.Sub(video.CreatedAt) <= time.Second)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testQueriesCreateVideo(t *testing.T) db.GetVideoByIDRow {
	user := testCreateUser(t)
	params := db.CreateVideoParams{
		UserID:   user.ID,
		Title:    utils.RandomString(10),
		PlayUrl:  utils.RandomString(20),
		CoverUrl: utils.RandomString(20),
	}
	t1 := time.Now()
	videoID, err := dao.Group.Mysql.CreateVideo(context.Background(), params)
	require.NoError(t, err)
	require.NotZero(t, videoID)
	video, err := dao.Group.Mysql.GetVideoByID(context.Background(), videoID)
	require.NoError(t, err)
	require.Equal(t, params.Title, video.Title)
	require.Equal(t, params.UserID, video.UserID)
	require.Equal(t, params.PlayUrl, video.PlayUrl)
	require.True(t, video.CreatedAt.Sub(t1) <= time.Second, t1.Sub(video.CreatedAt) <= time.Second)
	return video
}

func TestQueries_GetVideoByID(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				video := testQueriesCreateVideo(t)
				result, err := dao.Group.Mysql.GetVideoByID(context.Background(), video.ID)
				require.NoError(t, err)
				require.Equal(t, result.ID, video.ID)
				require.NoError(t, dao.Group.Mysql.DeleteVideo(context.Background(), result.ID))
				result, err = dao.Group.Mysql.GetVideoByID(context.Background(), video.ID)
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

func TestQueries_GetVideosByDate(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				time.Sleep(2 * time.Second)
				nums := int(utils.RandomInt(2, 10))
				resultMap := make(map[int64]bool, nums)
				videos := make([]db.GetVideoByIDRow, 0, nums)
				for i := 0; i < nums; i++ {
					video := testQueriesCreateVideo(t)
					videos = append(videos, video)
					resultMap[video.ID] = true
				}
				results, err := dao.Group.Mysql.GetVideosByDate(context.Background(), db.GetVideosByDateParams{
					CreatedAt: time.Now(),
					Limit:     int32(nums),
				})
				now := time.Now()
				require.NoError(t, err)
				require.Len(t, results, nums)
				for i := range results {
					require.True(t, resultMap[results[i].ID])
					require.True(t, !now.Before(results[i].CreatedAt))
					now = results[i].CreatedAt
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

func TestQueries_GetVideosByUserID(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				user := testCreateUser(t)
				nums := int(utils.RandomInt(1, 100))
				resultMap := make(map[int64]bool, nums)
				for i := 0; i < nums; i++ {
					videoID, err := dao.Group.Mysql.CreateVideo(context.Background(), db.CreateVideoParams{
						UserID:   user.ID,
						Title:    utils.RandomString(20),
						PlayUrl:  utils.RandomString(20),
						CoverUrl: utils.RandomString(20),
					})
					require.NoError(t, err)
					require.NotZero(t, videoID)
					resultMap[videoID] = true
				}
				results, err := dao.Group.Mysql.GetVideosByUserID(context.Background(), db.GetVideosByUserIDParams{
					UserID:   user.ID,
					FromID:   user.ID,
					UserID_2: user.ID,
				})
				require.NoError(t, err)
				require.NotNil(t, results)
				require.Len(t, results, nums)
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

func TestQueries_DeleteVideo(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				video := testQueriesCreateVideo(t)
				result, err := dao.Group.Mysql.GetVideoByID(context.Background(), video.ID)
				require.NoError(t, err)
				require.NoError(t, dao.Group.Mysql.DeleteVideo(context.Background(), result.ID))
				result, err = dao.Group.Mysql.GetVideoByID(context.Background(), video.ID)
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

func TestQueries_UpdateVideoCommentCount(t *testing.T) {
	type args struct {
		ctx context.Context
		arg db.UpdateVideoCommentCountParams
	}
	video := testQueriesCreateVideo(t)
	sum := video.CommentCount
	n1 := utils.RandomInt(1, 100)
	n2 := -utils.RandomInt(0, sum+n1)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "good: add num",
			args: args{ctx: context.Background(), arg: db.UpdateVideoCommentCountParams{
				CommentCount: n1,
				ID:           video.ID,
			}},
		},
		{
			name: "good: add -num",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateVideoCommentCountParams{
					CommentCount: n2,
					ID:           video.ID,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := dao.Group.Mysql.UpdateVideoCommentCount(tt.args.ctx, tt.args.arg)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				sum += tt.args.arg.CommentCount
				result, err := dao.Group.Mysql.GetVideoByID(context.Background(), tt.args.arg.ID)
				require.NoError(t, err)
				require.NotEmpty(t, result)
				require.Equal(t, result.CommentCount, sum)
			}
		})
	}
}

func TestQueries_UpdateVideoFavoriteCountCount(t *testing.T) {
	type args struct {
		ctx context.Context
		arg db.UpdateVideoFavoriteCountParams
	}
	video := testQueriesCreateVideo(t)
	sum := video.CommentCount
	n1 := utils.RandomInt(1, 100)
	n2 := -utils.RandomInt(0, sum+n1)
	n3 := -(sum + n1 + n2 + 1)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "good: add num",
			args: args{ctx: context.Background(), arg: db.UpdateVideoFavoriteCountParams{
				FavoriteCount: n1,
				ID:            video.ID,
			}},
		},
		{
			name: "good: add -num",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateVideoFavoriteCountParams{
					FavoriteCount: n2,
					ID:            video.ID,
				},
			},
		},
		{
			name: "bad: add < 0",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateVideoFavoriteCountParams{
					FavoriteCount: n3,
					ID:            video.ID,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := dao.Group.Mysql.UpdateVideoFavoriteCount(tt.args.ctx, tt.args.arg)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				sum += tt.args.arg.FavoriteCount
				result, err := dao.Group.Mysql.GetVideoByID(context.Background(), tt.args.arg.ID)
				require.NoError(t, err)
				require.NotEmpty(t, result)
				require.Equal(t, result.FavoriteCount, sum)
			}
		})
	}
}

func TestQueries_GetVideosByUserVideo(t *testing.T) {
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
					_, err := dao.Group.Mysql.CreateUserVideo(context.Background(), db.CreateUserVideoParams{
						UserID:  user.ID,
						VideoID: video.ID,
					})
					require.NoError(t, err)
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
