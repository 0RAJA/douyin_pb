package logic

import (
	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/model/reply"
	"github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type userFollow struct {
}

func (userFollow) AddRelation(ctx *gin.Context, params *request.RelationAction) errcode.Err {
	palLoad, err := global.Maker.VerifyToken(params.Token)
	if err != nil {
		return errcode.ErrServer
	}

	if params.ActionType == 1 {
		err := dao.Group.Mysql.CreateUserFollower(ctx, db.CreateUserFollowerParams{
			FromID: palLoad.UserID,
			ToID:   params.ToUserID,
		})
		if err != nil {
			return errcode.ErrServer
		}
		return nil
	}

	err1 := dao.Group.Mysql.DeleteUserFollower(ctx, db.DeleteUserFollowerParams{
		FromID: palLoad.UserID,
		ToID:   params.ToUserID,
	})
	if err1 != nil {
		return errcode.ErrServer
	}
	return nil
}

func (userFollow) GetUserFollowList(ctx *gin.Context, params *request.RelationFollowerList) (errcode.Err, reply.RelationFollowerList) {
	userFollowers, err := dao.Group.Mysql.GetUserFollowersByFromUserID(ctx, params.UserID)
	if err != nil {
		return errcode.ErrServer, reply.RelationFollowerList{}
	}
	users := make([]db.User, 0)
	for _, follower := range userFollowers {
		user, err := dao.Group.Mysql.GetUserByUserID(ctx, follower.ToID)
		if err != nil {
			return nil, reply.RelationFollowerList{}
		}
		users = append(users, user)
	}

	return nil, reply.RelationFollowerList{
		UserList: users,
	}
}
