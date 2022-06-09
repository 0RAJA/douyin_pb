package v1

import (
	"time"

	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/logic"
	mid "github.com/0RAJA/douyin/internal/middleware"
	"github.com/0RAJA/douyin/internal/model/reply"
	"github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type video struct {
}

func (video) Feed(c *gin.Context) {
	rly := app.NewReply(c)
	params := request.Feed{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.SendErr(reply.Feed{}, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.SendErr(reply.Feed{}, err)
		return
	}
	var userID int64
	payload, err := mid.GetPayload(c)
	if err != nil {
		userID = 0
	} else {
		userID = payload.UserID
	}
	results, err := logic.Group.Video.Feed(c, db.GetVideosByDateParams{
		UserID:    userID,
		FromID:    userID,
		CreatedAt: time.Unix(params.LatestTime, 0),
		Limit:     30,
	})
	if err != nil {
		rly.SendErr(reply.Feed{}, err)
		return
	}
	rly.SendData(results)
}

func (video) PublishAction(c *gin.Context) {
	rly := app.NewReply(c)
	params := request.PublishAction{}
	if err := c.ShouldBind(&params); err != nil {
		rly.SendErr(nil, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := logic.Group.Video.PublishAction(c, params.Data, params.Title); err != nil {
		rly.SendErr(nil, err)
		return
	}
	rly.SendData(nil)
}

func (video) PublishList(c *gin.Context) {
	rly := app.NewReply(c)
	params := request.PublishList{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.SendErr(nil, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.Video.PublishList(c, params.UserID)
	if err != nil {
		rly.SendErr(nil, err)
		return
	}
	rly.SendData(result)
}
