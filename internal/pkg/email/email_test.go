//go:build mail
// +build mail

package email

import (
	"fmt"
	"testing"
	"time"

	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/times"
	"github.com/stretchr/testify/require"
)

func TestEmailSendMail(t *testing.T) {
	defailtMailer := NewEmail(&SMTPInfo{
		Host:     global.Settings.Email.Host,
		Port:     global.Settings.Email.Port,
		IsSSL:    global.Settings.Email.IsSSL,
		UserName: global.Settings.Email.UserName,
		Password: global.Settings.Email.Password,
		From:     global.Settings.Email.From,
	})
	err := defailtMailer.SendMail( // 短信通知
		global.Settings.Email.To,
		fmt.Sprintf("异常抛出，发生时间: %s,%d", times.GetNowDateTimeStr(), time.Now().Unix()),
		fmt.Sprintf("错误信息: %v", "test"),
	)
	require.NoError(t, err)
}
