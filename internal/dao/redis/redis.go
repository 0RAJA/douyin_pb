package redis

import (
	"context"

	"github.com/0RAJA/douyin/internal/global"
	rds "github.com/go-redis/redis/v8"
)

type Redis struct {
	rdb *rds.Client
}

func Init(Addr, Password string, PoolSize, DB int) Redis {
	rdb := rds.NewClient(&rds.Options{
		Addr:     Addr,     // ip:端口
		Password: Password, // 密码
		PoolSize: PoolSize, // 连接池
		DB:       DB,       // 默认连接数据库
	})
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Server.DefaultContextTimeout)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil { // 测试连接
		panic(err)
	}
	return Redis{rdb: rdb}
}
