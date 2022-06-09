package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/routing/router"
	"github.com/0RAJA/douyin/internal/setting"
	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

func initSettings() {
	setting.Group.Config.Init()
	setting.Group.Dao.Init()
	setting.Group.Snowflake.Init()
	setting.Group.Maker.Init()
	setting.Group.Log.Init()
}

func main() {
	initSettings() // 初始化配置文件
	if global.Settings.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.NewRouter() // 注册路由
	s := &http.Server{
		Addr:           global.Settings.Server.Address,
		Handler:        r,
		ReadTimeout:    global.Settings.Server.ReadTimeout,
		WriteTimeout:   global.Settings.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("Server started!")
	fmt.Println("AppName:", global.Settings.App.Name, "Version:", global.Settings.App.Version, "Address:", global.Settings.Server.Address, "RunMode:", global.Settings.Server.RunMode)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			global.Logger.Info(err.Error())
		}
	}()
	gracefulExit(s) // 优雅退出
	global.Logger.Info("Server exited!")
}

// 优雅退出
func gracefulExit(s *http.Server) {
	// 退出通知
	quit := make(chan os.Signal, 1)
	// 等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("ShutDown Server...")
	// 给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Server.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { // 优雅退出
		global.Logger.Info("Server forced to ShutDown,Err:" + err.Error())
	}
}
