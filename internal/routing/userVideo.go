package routing

import (
	"github.com/gin-gonic/gin"
)

type userVideo struct {
}

func (u *userVideo) Init(Group *gin.RouterGroup) {
	routerGroup := Group.Group("/favorite")
	{
		routerGroup.POST("/action")
	}
}
