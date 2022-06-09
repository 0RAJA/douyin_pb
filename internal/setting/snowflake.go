package setting

import (
	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/snowflake"
)

type sf struct {
}

// Init 雪花算法初始化
func (sf) Init() {
	var err error
	if global.Snowflake, err = snowflake.Init(global.Settings.App.StartTime, global.Settings.App.Format, 1); err != nil {
		panic(err)
	}
}
