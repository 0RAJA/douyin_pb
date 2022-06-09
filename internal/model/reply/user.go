package reply

import (
	"github.com/0RAJA/douyin/internal/model/common"
)

type Auth struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegister struct {
	Auth
}

type UserLogin struct {
	Auth
}

type User struct {
	common.User
	IsFollow bool `json:"is_follow"`
}

type UserInfo struct {
	User User `json:"user"`
}
