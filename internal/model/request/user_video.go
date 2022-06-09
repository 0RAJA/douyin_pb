package request

import (
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
)

type FavoriteAction struct {
	Token      string `form:"token" binding:"required"`
	VideoID    int64  `form:"video_id" binding:"required,gte=1"`
	ActionType uint8  `form:"action_type" binding:"required"`
}

func (fa *FavoriteAction) Judge() errcode.Err {
	switch fa.ActionType {
	case 1, 2:
	default:
		return errcode.ErrParamsNotValid
	}
	return nil
}
