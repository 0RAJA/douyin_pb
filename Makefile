install_golang-cli: # 安装golang-cli工具，用于静态检查代码质量
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
test: # 静态检查代码质量以及运行所有的测试程序
	golangci-lint run && go test -v -cover ./...
mysql_init: # mysql初始化
	docker run --name mysql_80 --privileged=true -p 3366:3306 -v mysql_80_data:/var/lib/mysql -v mysql_80_conf:/etc/mysql -e MYSQL_ROOT_PASSWORD='123456' -e MYSQL_DATABASE=douyin -e LANG=C.UTF-8 -e character_set_database=utf8 -d mysql:8.0
redis_init: # redis初始化
	docker run --name redis_62 --privileged=true -p 7963:7963 -v $(pwd)/config/redis:/redis -d redis:6.2 redis-server /redis/redis.conf
migrate_install: # 安装migrate数据库迁移工具
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz & sudo mv migrate /usr/bin/migrate
migrate_init_db: # 初始化迁移文件
	migrate create -ext sql -dir internal/dao/mysql/migration -seq init_schema
migrate_up: # 构建迁移工具并且进行数据库迁移
	go generate -x ./... && go build -o bin/migrate cmd/migrate/migrate.go && bin/migrate -addr=localhost:3366 -source=internal/dao/mysql/migration -dbName=douyin -auth=root:123456
sqlc: # sqlc生成go代码
	bin/sqlc generate
run: # 本机启动程序
	go generate -x ./... && go mod tidy && go run cmd/main/main.go
build: # 本机构建程序
	go generate -x ./... && go build -o bin/main cmd/main/main.go
docker_build: # 构建镜像
	docker build -t douyin_douyin:latest .
docker_connect_net: # 创建docker网络
	docker network create douyin && docker network connect douyin mysql_80 && docker network connect douyin redis_62
docker_run: # 启动镜像
	docker run --name douyin -p 8080:8080 --net douyin -d douyin_douyin:latest
pull:
	git fetch origin master && git rebase origin/master
goimports_install: # goimports安装
	go get golang.org/x/tools/cmd/goimports
