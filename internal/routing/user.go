package routing

import (
	v1 "github.com/0RAJA/douyin/internal/api/v1"
	mid "github.com/0RAJA/douyin/internal/middleware"
	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user *user) Init(group *gin.RouterGroup) {
	userGroup := group.Group("user")
	{
		userGroup.POST("register/", v1.Group.User.Register)
		userGroup.POST("login/", v1.Group.User.Login)
		userGroup.GET("/", mid.MustLogin, v1.Group.User.UserInfo)
	}
}
