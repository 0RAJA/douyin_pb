package email

import (
	"os"
	"testing"

	"github.com/0RAJA/douyin/internal/setting"
)

func TestMain(m *testing.M) {
	setting.Group.Config.Init()
	os.Exit(m.Run())
}
