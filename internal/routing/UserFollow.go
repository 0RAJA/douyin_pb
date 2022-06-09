package routing

import (
	v1 "github.com/0RAJA/douyin/internal/api/v1"
	"github.com/gin-gonic/gin"
)

type userFollow struct {
}

func (u *userFollow) Init(Group *gin.RouterGroup) {
	relationGroup := Group.Group("relation")
	{
		relationGroup.POST("/action", v1.Group.UserFollow.DoAttention)
		relationGroup.GET("/follow/list", v1.Group.UserFollow.GetUserFollowList)
	}
}
