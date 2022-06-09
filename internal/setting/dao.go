package setting

import (
	"github.com/0RAJA/douyin/internal/dao"
	"github.com/0RAJA/douyin/internal/dao/mysql"
	"github.com/0RAJA/douyin/internal/dao/redis"
	"github.com/0RAJA/douyin/internal/global"
)

type mDao struct {
}

// Init 持久化层初始化
func (m mDao) Init() {
	dao.Group.Mysql = mysql.Init(global.Settings.Mysql.DriverName, global.Settings.Mysql.SourceName)
	dao.Group.Redis = redis.Init(global.Settings.Redis.Address, global.Settings.Redis.Password, global.Settings.Redis.PoolSize, global.Settings.Redis.DB)
}
