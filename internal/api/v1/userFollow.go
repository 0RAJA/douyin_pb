package v1

import (
	"github.com/0RAJA/douyin/internal/logic"
	"github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type userFollow struct {
}

func (userFollow) DoAttention(ctx *gin.Context) {
	rly := app.NewReply(ctx)
	var req *request.RelationAction

	if err := ctx.ShouldBindQuery(req); err != nil {
		rly.SendErr(nil, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}

	if err := req.Judge(); err != nil {
		rly.SendErr(nil, errcode.ErrParamsNotValid)
		return
	}

	err := logic.Group.UserFollow.AddRelation(ctx, req)
	if err != nil {
		rly.SendErr(nil, err)
		return
	}

	rly.SendData(nil)
}

func (userFollow) GetUserFollowList(ctx *gin.Context) {
	rly := app.NewReply(ctx)
	var req *request.RelationFollowerList

	if err := ctx.ShouldBindQuery(req); err != nil {
		rly.SendErr(nil, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}

	err, data := logic.Group.UserFollow.GetUserFollowList(ctx, req)
	if err != nil {
		rly.SendErr(nil, err)
	}
	rly.SendData(data)
}
