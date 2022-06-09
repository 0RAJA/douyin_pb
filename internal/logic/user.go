package logic

import (
	"database/sql"
	"errors"

	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/global"
	mid "github.com/0RAJA/douyin/internal/middleware"
	"github.com/0RAJA/douyin/internal/model/common"
	"github.com/0RAJA/douyin/internal/model/reply"
	"github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/0RAJA/douyin/internal/pkg/password"
	"github.com/gin-gonic/gin"
)

type user struct{}

func (user) Register(ctx *gin.Context, params request.UserRegister) (reply.UserRegister, errcode.Err) {
	if _, err := dao.Group.Mysql.GetUserByUsername(ctx, params.Username); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			global.Logger.Error(err.Error())
			return reply.UserRegister{}, errcode.ErrServer
		}
	} else {
		return reply.UserRegister{}, errcode.ErrUserHasExist
	}
	hashPasswd, err := password.HashPassword(params.Password)
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.UserRegister{}, errcode.ErrServer
	}
	userID, err := dao.Group.Mysql.CreateUser(ctx, db.CreateUserParams{
		Name:     params.Username,
		Password: hashPasswd,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.UserRegister{}, errcode.ErrServer
	}
	token, err := global.Maker.CreateToken(userID, params.Username, global.Settings.Token.AssessTokenDuration)
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.UserRegister{}, errcode.ErrServer
	}
	return reply.UserRegister{Auth: reply.Auth{
		UserID: userID,
		Token:  token,
	}}, nil
}

func (user) Login(ctx *gin.Context, params request.UserLogin) (reply.UserLogin, errcode.Err) {
	user, err := dao.Group.Mysql.GetUserByUsername(ctx, params.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			global.Logger.Info(err.Error())
			return reply.UserLogin{}, errcode.ErrUserNotExist
		}
		global.Logger.Error(err.Error())
		return reply.UserLogin{}, errcode.ErrServer
	}
	if err := password.CheckPassword(params.Password, user.Password); err != nil {
		global.Logger.Info(err.Error())
		return reply.UserLogin{}, errcode.ErrNamePasswordNotMatch
	}
	token, err := global.Maker.CreateToken(user.ID, params.Username, global.Settings.Token.AssessTokenDuration)
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.UserLogin{}, errcode.ErrServer
	}
	return reply.UserLogin{Auth: reply.Auth{
		UserID: user.ID,
		Token:  token,
	}}, nil
}

func (user) UserInfo(ctx *gin.Context, params request.UserInfo) (reply.UserInfo, errcode.Err) {
	payload, err := mid.GetPayload(ctx)
	if err != nil {
		return reply.UserInfo{}, err
	}
	userInfo, err1 := dao.Group.Mysql.GetUserByUserID(ctx, params.UserID)
	if err1 != nil {
		if errors.Is(err1, sql.ErrNoRows) {
			global.Logger.Info(err1.Error())
			return reply.UserInfo{}, errcode.ErrNotFound
		}
		global.Logger.Error(err1.Error())
		return reply.UserInfo{}, errcode.ErrServer
	}
	isFollow := false
	if _, err := dao.Group.Mysql.GetUserFollower(ctx, db.GetUserFollowerParams{
		FromID: payload.UserID,
		ToID:   params.UserID,
	}); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			global.Logger.Error(err.Error())
			return reply.UserInfo{}, errcode.ErrServer
		}
	} else {
		isFollow = true
	}
	return reply.UserInfo{User: reply.User{
		User: common.User{
			ID:            userInfo.ID,
			Name:          userInfo.Name,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
		},
		IsFollow: isFollow,
	}}, nil
}
