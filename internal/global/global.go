package global

import (
	"github.com/0RAJA/douyin/internal/model/config"
	"github.com/0RAJA/douyin/internal/pkg/logger"
	"github.com/0RAJA/douyin/internal/pkg/snowflake"
	"github.com/0RAJA/douyin/internal/pkg/token"
)

var (
	Logger    *logger.Log          // 日志
	Settings  config.All           // 全局配置
	Maker     token.Maker          // 操作token
	Snowflake *snowflake.Snowflake // 生成ID
	RootDir   string               // 项目跟路径
)
