package mysql

import (
	"context"
	"database/sql"

	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/global"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql interface {
	db.Store
}

func Init(driverName, dataSourceName string) Mysql {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Server.DefaultContextTimeout)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		panic(err)
	}
	return &db.SqlStore{Queries: db.New(conn), DB: conn}
}
