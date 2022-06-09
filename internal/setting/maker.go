package setting

import (
	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/token"
)

type maker struct {
}

// Init tokenMaker初始化
func (maker) Init() {
	var err error
	global.Maker, err = token.NewPasetoMaker([]byte(global.Settings.Token.Key))
	if err != nil {
		panic(err)
	}
}
