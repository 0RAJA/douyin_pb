package setting

import (
	"flag"
	"strings"

	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/setting"
)

type config struct {
}

var (
	configPaths string // 配置文件路径
	configName  string // 配置文件名
	configType  string // 配置文件类型
)

func setupFlag() {
	// 命令行参数绑定
	flag.StringVar(&configName, "name", "app", "配置文件名")
	flag.StringVar(&configType, "type", "yml", "配置文件类型")
	flag.StringVar(&configPaths, "path", global.RootDir+"/config/app", "指定要使用的配置文件路径")
	flag.Parse()
}

// Init 读取配置文件
func (config) Init() {
	setupFlag()
	newSetting, err := setting.NewSetting(configName, configType, strings.Split(configPaths, ",")...) // 引入配置文件路径
	if err != nil {
		panic(err)
	}
	if err := newSetting.BindAll(&global.Settings); err != nil {
		panic(err)
	}
}
