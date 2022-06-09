package dao

import (
	"github.com/0RAJA/douyin/internal/dao/mysql"
	"github.com/0RAJA/douyin/internal/dao/redis"
)

type group struct {
	Mysql mysql.Mysql
	Redis redis.Redis
}

var Group = new(group)
