package request

import (
	"fmt"

	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
)

type CommentAction struct {
	Token       string `form:"token" binding:"required"`
	VideoID     int64  `form:"video_id" binding:"required,gte=1"`
	ActionType  uint8  `form:"action_type" binding:"required"`
	CommentID   int64  `form:"comment_id"`
	CommentText string `form:"comment_text"`
}

func (opt *CommentAction) Judge() (err errcode.Err) {
	switch opt.ActionType {
	case 1:
		if now := len(opt.CommentText); now > global.Settings.Rule.CommentLenMax || now < global.Settings.Rule.CommentLenMin {
			err = errcode.ErrLength.WithDetails(fmt.Sprintf("comment_text length' maximum is:%d,minimum is:%d", global.Settings.Rule.CommentLenMax, global.Settings.Rule.CommentLenMin))
		}
	case 2:
		if opt.CommentID <= 0 {
			err = errcode.ErrParamsNotValid
		}
	default:
		err = errcode.ErrParamsNotValid
	}
	return
}

type CommentList struct {
	VideoID int64 `form:"video_id" binding:"required,gte=1"`
}
