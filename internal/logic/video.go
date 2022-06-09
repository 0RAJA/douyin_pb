package logic

import (
	"context"
	"mime/multipart"

	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/global"
	mid "github.com/0RAJA/douyin/internal/middleware"
	"github.com/0RAJA/douyin/internal/model/common"
	"github.com/0RAJA/douyin/internal/model/reply"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/0RAJA/douyin/internal/upload"
	"github.com/gin-gonic/gin"
)

type video struct {
}

func (video) Feed(ctx context.Context, params db.GetVideosByDateParams) (reply.Feed, errcode.Err) {
	videos, err := dao.Group.Mysql.GetVideosByDate(ctx, params)
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.Feed{}, errcode.ErrServer
	}
	results := make([]reply.Video, len(videos))
	for i := range videos {
		results[i] = reply.Video{
			ID: videos[i].ID,
			Author: reply.User{
				User: common.User{
					ID:            videos[i].UserID,
					Name:          videos[i].Name,
					FollowCount:   videos[i].FollowCount,
					FollowerCount: videos[i].FollowerCount,
				},
				IsFollow: videos[i].IsFavoriteUser,
			},
			Title:         videos[i].Title,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    videos[i].IsFavoriteVideo,
		}
	}
	next := int64(0)
	if len(results) > 0 {
		next = videos[len(videos)-1].CreatedAt.Unix()
	}
	return reply.Feed{
		NextTime:  next,
		VideoList: results,
	}, nil
}

func (video) PublishAction(c *gin.Context, data *multipart.FileHeader, title string) errcode.Err {
	oss := upload.NewOSS()
	paload, mErr := mid.GetPayload(c)
	if mErr != nil {
		return mErr
	}
	url, _, err := oss.UploadFile(data)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ErrServer
	}
	_, err = dao.Group.Mysql.CreateVideo(c, db.CreateVideoParams{
		UserID:   paload.UserID,
		Title:    title,
		PlayUrl:  url,
		CoverUrl: global.Settings.Rule.DefaultCoverURL,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ErrServer
	}
	return nil
}

func (video) PublishList(c *gin.Context, userID int64) (reply.PublishList, errcode.Err) {
	videos, err := dao.Group.Mysql.GetVideosByUserID(c, db.GetVideosByUserIDParams{
		UserID:   userID,
		FromID:   userID,
		UserID_2: userID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.PublishList{}, errcode.ErrServer
	}
	results := make([]reply.Video, len(videos))
	for i := range videos {
		results[i] = reply.Video{
			ID: videos[i].ID,
			Author: reply.User{
				User: common.User{
					ID:            videos[i].UserID,
					Name:          videos[i].Name,
					FollowCount:   videos[i].FollowCount,
					FollowerCount: videos[i].FollowerCount,
				},
				IsFollow: videos[i].IsFavoriteUser,
			},
			Title:         videos[i].Title,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    videos[i].IsFavoriteVideo,
		}
	}
	return reply.PublishList{VideoList: results}, nil
}
