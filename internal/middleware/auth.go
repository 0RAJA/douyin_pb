package middleware

import (
	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/app"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/0RAJA/douyin/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationKey = "payload"
	TokenKey         = "token"
)

func GetPayload(ctx *gin.Context) (*token.Payload, errcode.Err) {
	payload, ok := ctx.Get(AuthorizationKey)
	if !ok {
		return nil, errcode.ErrInsufficientPermissions
	}
	return payload.(*token.Payload), nil
}

// MustLogin 必须登陆
func MustLogin(ctx *gin.Context) {
	reply := app.NewReply(ctx)
	if _, ok := ctx.Get(AuthorizationKey); !ok {
		reply.SendErr(nil, errcode.ErrUnauthorizedAuthNotExist)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// Auth 默认鉴权
func Auth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var inToken string
		inToken = ctx.Query(TokenKey)
		if inToken == "" {
			inToken = ctx.PostForm(TokenKey)
		}
		payload, err := global.Maker.VerifyToken(inToken)
		if err != nil {
			ctx.Next()
			return
		}
		ctx.Set(AuthorizationKey, payload)
		ctx.Next()
	}
}
