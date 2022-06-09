package routing

import (
	v1 "github.com/0RAJA/douyin/internal/api/v1"
	"github.com/gin-gonic/gin"
)

type comment struct {
}

// comment註冊

func (com *comment) Init(group *gin.RouterGroup) {
	commentGroup := group.Group("comment")
	{
		commentGroup.POST("/action", v1.Group.Comment.CommentAction)

		commentGroup.GET("list", v1.Group.Comment.List)
	}
}
