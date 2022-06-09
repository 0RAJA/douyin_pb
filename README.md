# 抖音项目

构建环境:`Linux x86_64`

语言版本:`go version go1.16.14 linux/amd64`

项目结构说明

```bash

   .
├── bin # 二进制可执行文件目录
│   ├── main # 主程序可执行文件
│   └── migrate # 数据库迁移的可执行文件
├── cmd # 入口函数所在目录
│   ├── main # 主程序
│   │   └── main.go
│   └── migrate # 数据库迁移程序
│       └── migrate.go
├── config # 配置文件
│   ├── app # 本地配置文件(默认)
│   ├── app_docker # docker配置文件(以docker-compose运行默认此配置)
│   └── redis # redis配置文件
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal # 主代码目录
│   ├── api # 接口处理路径
│   │   └── v1 # v1 版本接口目录
│   │       ├── enter.go # api 入口点
│   │       └── user.go # user的api
│   ├── dao # 持久化层
│   │   ├── enter.go # dao入口点
│   │   ├── mysql
│   │   │   ├── migration # 表结构，用于数据库迁移
│   │   │   │   ├── 000001_init_schema.down.sql # 删除结构
│   │   │   │   └── 000001_init_schema.up.sql # 创建数据库结构
│   │   │   ├── mysql.go
│   │   │   ├── query # 查询语句，使用sqlc生成go代码
│   │   │   │   ├── user_followers.sql # user_followers 相关操作
│   │   │   │   └── user.sql # user 相关操作
│   │   │   └── sqlc # sqlc生成的代码目录 每个操作需要详细的单元测试
│   │   └── redis
│   │       └── redis.go
│   ├── global # 全局变量
│   │   ├── global.go # 全局
│   │   ├── infer.go # 推断root
│   │   └── infer_test.go
│   ├── logic # 逻辑处理层
│   │   ├── enter.go # logic 入口点
│   │   └── user.go # user logic
│   ├── middleware # 中间件
│   │   ├── auth.go # 鉴权
│   │   ├── cores.go # 跨域
│   │   ├── limiter.go # 限流
│   │   ├── logger.go # 日志
│   │   └── recovery.go # 恢复
│   ├── model # 模型目录
│   │   ├── common # 通用模型
│   │   │   └── common.go
│   │   ├── config # 配置文件模型，用户绑定配置文件
│   │   │   └── config.go
│   │   ├── reply # 回复模型
│   │   │   ├── common.go
│   │   │   └── user.go
│   │   └── request # 请求模型
│   │       └── user.go
│   ├── pkg # 相关组件
│   │   ├── app # 用于格式的规范
│   │   │   ├── errcode
│   │   │   │   ├── codes.go # 所有错误码放这里
│   │   │   │   └── err.go
│   │   │   └── reply.go # 回复
│   │   ├── email # 邮箱
│   │   ├── limiter # 限流
│   │   ├── logger # 日志
│   │   ├── password # 密码加密
│   │   ├── setting # 配置文件读取
│   │   ├── snowflake # 生成ID
│   │   ├── times # 时间包
│   │   ├── token # 生成token
│   │   ├── upload # oss
│   │   └── utils # 工具包
│   ├── routing
│   │   ├── enter.go # 路由层入口
│   │   ├── router
│   │   │   └── router.go # 新建路由
│   │   └── user.go
│   └── setting # 初始化相关流程
└── storage # 相关数据持久化
    └── Applogs # 产生的日志
├── Makefile # 便捷操作
├── README.md
├── sqlc.yaml # sqlc配置文件
├── start.sh # 用于构建时初始化数据库
├── storage
│   └── Applogs
│       ├── error.log # 错误日志
│       └── info.log # 常规日志
└── wait-for.sh # 用于构建时服务间同步
```
# 项目docker运行
`docker-compose up`
# 项目本机测试
```
# 初始化 mysql 本机端口3366
make mysql_init
# 初始化 redis 运行不了就把对应makefile执行语句复制到终端执行
make redis_init 本机端口7963
然后启动本地项目就行
```
