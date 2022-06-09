package v1

import (
	"github.com/0RAJA/douyin/internal/logic"
	"github.com/0RAJA/douyin/internal/model/reply"
	request2 "github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type comment struct {
}

//評論操作
func (comment) CommentAction(c *gin.Context) {
	reply := app.NewReply(c)
	var req *request2.CommentAction
	if err := c.ShouldBindQuery(req); err != nil {
		reply.SendErr(nil, errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}

	err, data := logic.Group.Comment.AddComment(c, req)

	if err != nil {
		reply.SendErr(nil, errcode.ErrServer)
	}

	reply.SendData(data)
}

func (comment) List(ctx *gin.Context) {
	rly := app.NewReply(ctx)
	params := request2.CommentList{}
	var result reply.CommentList
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rly.SendErr(result, errcode.ErrParamsNotValid.WithDetails(err.Error()))
	}

	result, err := logic.Group.Comment.List(ctx, params)
	if err != nil {
		rly.SendErr(result, err)
		return
	}
	rly.SendData(result)

}
