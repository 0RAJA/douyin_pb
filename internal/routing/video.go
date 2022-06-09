package routing

import (
	v1 "github.com/0RAJA/douyin/internal/api/v1"
	"github.com/gin-gonic/gin"
)

type video struct {
}

func (video) Init(r *gin.RouterGroup) {
	r.GET("feed", v1.Group.Video.Feed)
	r.POST("publish/action/", v1.Group.Video.PublishAction)
	r.GET("publish/list/", v1.Group.Video.PublishList)
}
