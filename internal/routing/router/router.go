package router

import (
	mid "github.com/0RAJA/douyin/internal/middleware"
	"github.com/0RAJA/douyin/internal/routing"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors(), mid.GinRecovery(true), mid.GinLogger(), mid.Auth())
	root := r.Group("douyin")
	{
		routing.Group.User.Init(root)
		routing.Group.Video.Init(root)
		routing.Group.UserFollow.Init(root)
		routing.Group.UserVideo.Init(root)
		routing.Group.Comment.Init(root)
	}
	return r
}
