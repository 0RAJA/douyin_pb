package v1

import (
	"github.com/0RAJA/douyin/internal/logic"
	"github.com/0RAJA/douyin/internal/model/reply"
	"github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type user struct{}

func (user) Login(ctx *gin.Context) {
	rly := app.NewReply(ctx)
	params := request.UserLogin{}
	var result reply.UserLogin
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rly.SendErr(result, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.SendErr(result, err)
		return
	}
	result, err := logic.Group.User.Login(ctx, params)
	if err != nil {
		rly.SendErr(result, err)
		return
	}
	rly.SendData(result)
}

func (user) Register(ctx *gin.Context) {
	rly := app.NewReply(ctx)
	params := request.UserRegister{}
	var result reply.UserRegister
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rly.SendErr(result, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := params.Judge(); err != nil {
		rly.SendErr(result, err)
		return
	}
	result, err := logic.Group.User.Register(ctx, params)
	if err != nil {
		rly.SendErr(result, err)
		return
	}
	rly.SendData(result)
}

func (user) UserInfo(ctx *gin.Context) {
	rly := app.NewReply(ctx)
	params := request.UserInfo{}
	var result reply.UserInfo
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rly.SendErr(result, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.User.UserInfo(ctx, params)
	if err != nil {
		rly.SendErr(result, err)
		return
	}
	rly.SendData(result)
}
