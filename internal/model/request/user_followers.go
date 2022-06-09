package request

import (
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
)

type RelationAction struct {
	Token      string `form:"token"`
	ToUserID   int64  `form:"to_user_id" binding:"required,gte=1"`
	ActionType uint8  `form:"action_type" binding:"required"`
}

func (ra *RelationAction) Judge() errcode.Err {
	switch ra.ActionType {
	case 1, 2:
	default:
		return errcode.ErrParamsNotValid
	}
	return nil
}

type RelationFollowList struct {
	UserID int64 `form:"user_id" binding:"required,gte=1"`
}

type RelationFollowerList struct {
	Token  string `from:"token"`
	UserID int64  `form:"user_id" binding:"required,gte=1"`
}
