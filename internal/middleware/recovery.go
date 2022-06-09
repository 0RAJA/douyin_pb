package middleware

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/email"
	"github.com/0RAJA/douyin/internal/pkg/times"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 异常捕获处理

// 自定义 Recovery

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.Settings.Email.Host,
		Port:     global.Settings.Email.Port,
		IsSSL:    global.Settings.Email.IsSSL,
		UserName: global.Settings.Email.UserName,
		Password: global.Settings.Email.Password,
		From:     global.Settings.Email.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				err1 := defailtMailer.SendMail( // 短信通知
					global.Settings.Email.To,
					fmt.Sprintf("异常抛出，发生时间: %v", time.Now().Format(times.LayoutDateTime)),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err1 != nil {
					global.Logger.Error(fmt.Sprintf("mail.SendMail Error: %v", err1.Error()))
				}

				// Check for a broken connection
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.Logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if stack {
					global.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
