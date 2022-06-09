package request

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
)

type Feed struct {
	LatestTime int64 `form:"latest_time"`
}

func (feed *Feed) Judge() errcode.Err {
	if feed.LatestTime == 0 {
		feed.LatestTime = time.Now().Unix()
	}
	return nil
}

type PublishAction struct {
	Data  *multipart.FileHeader `form:"data" binding:"required"`
	Title string                `form:"title" binding:"required,gte=1"`
	Token string                `form:"token" binding:"required"`
}

func (pa *PublishAction) Judge() errcode.Err {
	if now := len(pa.Title); now > global.Settings.Rule.TitlesLenMax || now < global.Settings.Rule.TitlesLenMin {
		return errcode.ErrLength.WithDetails(fmt.Sprintf("title length's maximum is:%d,minimum is:%d", global.Settings.Rule.TitlesLenMax, global.Settings.Rule.TitlesLenMin))
	}
	return nil
}

type PublishList struct {
	UserID int64 `form:"user_id" binding:"required,gte=1"`
}
