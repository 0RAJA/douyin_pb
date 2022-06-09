package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// 用于数据库的迁移

var (
	sourceURL string
	address   string
	dbName    string
	auth      string
)

// 绑定相关参数
func main() {
	flag.StringVar(&sourceURL, "source", "migration", "migration文件夹路径")
	flag.StringVar(&address, "addr", "mysql_80:3306", "mysql address")
	flag.StringVar(&dbName, "dbName", "douyin", "数据库名")
	flag.StringVar(&auth, "auth", "root:123456", "用户名:密码")
	flag.Parse()
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s)/%s?multiStatements=true", auth, address, dbName))
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", sourceURL),
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		log.Println(err)
	}
	log.Println("db ok")
}
