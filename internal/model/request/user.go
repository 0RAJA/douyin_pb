package request

import (
	"fmt"

	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
)

type auth struct {
	Username string `form:"username" binding:"required,gte=1"`
	Password string `form:"password" binding:"required,gte=1"`
}

type UserRegister struct {
	auth
}

func (register *UserRegister) Judge() errcode.Err {
	if now := len(register.Username); now > global.Settings.Rule.UsernameLenMax || now < global.Settings.Rule.UsernameLenMin {
		return errcode.ErrLength.WithDetails(fmt.Sprintf("username length's maximum is:%d,minimum is:%d", global.Settings.Rule.UsernameLenMax, global.Settings.Rule.UsernameLenMin))
	}
	if now := len(register.Password); now > global.Settings.Rule.PasswordLenMax || now < global.Settings.Rule.PasswordLenMin {
		return errcode.ErrLength.WithDetails(fmt.Sprintf("password length's maximum is:%d,minimum is:%d", global.Settings.Rule.PasswordLenMax, global.Settings.Rule.PasswordLenMin))
	}
	return nil
}

type UserLogin struct {
	auth
}

func (login *UserLogin) Judge() errcode.Err {
	return (*UserRegister)(login).Judge()
}

type UserInfo struct {
	UserID int64 `form:"user_id" binding:"required,gte=1"`
}
