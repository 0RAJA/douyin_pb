package logic

import (
	"database/sql"
	"errors"

	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/model/common"
	"github.com/0RAJA/douyin/internal/model/reply"
	request2 "github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type comment struct {
}

func (comment) AddComment(ctx *gin.Context, params *request2.CommentAction) (errcode.Err, *reply.CommentAction) {
	palLoad, err2 := global.Maker.VerifyToken(params.Token)
	if err2 != nil {
		return errcode.ErrServer, nil
	}
	if params.ActionType == 1 {
		commentParams := db.CreateCommentParams{
			UserID:  palLoad.UserID,
			VideoID: params.VideoID,
			Content: params.CommentText,
		}
		comment, err3 := dao.Group.Mysql.CreateCommentWithTx(ctx, commentParams)
		if err3 != nil {
			return errcode.ErrServer, nil
		}

		u := common.User{
			ID:            comment.UserID,
			Name:          comment.Name,
			Password:      "",
			FollowCount:   comment.FollowCount,
			FollowerCount: comment.FollowerCount,
		}
		video, err := dao.Group.Mysql.GetVideoByID(ctx, params.VideoID)
		if err != nil {
			return errcode.ErrServer, nil
		}
		relations, err := dao.Group.Mysql.IsExistRelations(ctx, db.IsExistRelationsParams{
			FromID: comment.UserID,
			ToID:   video.UserID,
		})
		if err != nil {
			return errcode.ErrServer, nil
		}
		var isFollow bool
		if relations > 0 {
			isFollow = true
		} else {
			isFollow = false
		}
		u2 := reply.User{
			User:     u,
			IsFollow: isFollow,
		}
		data := &reply.CommentAction{
			ID:         comment.ID,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
			User:       u2,
		}
		return nil, data
	}

	err := dao.Group.Mysql.DeleteComment(ctx, params.CommentID)
	if err != nil {
		return errcode.ErrServer, nil
	}

	return nil, nil
}

func (comment) List(ctx *gin.Context, params request2.CommentList) (reply.CommentList, errcode.Err) {
	result := reply.CommentList{}
	com := reply.CommentAction{}
	re, err := dao.Group.Mysql.GetCommentsByVideoId(ctx, params.VideoID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			global.Logger.Info(err.Error())
			return reply.CommentList{}, errcode.ErrNotFound
		}
		global.Logger.Error(err.Error())
		return reply.CommentList{}, errcode.ErrServer
	}
	v, err := dao.Group.Mysql.GetVideoByID(ctx, params.VideoID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			global.Logger.Info(err.Error())
			return reply.CommentList{}, errcode.ErrNotFound
		}
		global.Logger.Error(err.Error())
		return reply.CommentList{}, errcode.ErrServer
	}
	userID := v.UserID
	for _, v := range re {
		com.ID = v.ID
		com.Content = v.Content
		com.CreateDate = v.CreateDate
		user, err := dao.Group.Mysql.GetUserByUserID(ctx, v.UserID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				global.Logger.Info(err.Error())
				return reply.CommentList{}, errcode.ErrNotFound
			}
			global.Logger.Error(err.Error())
			return reply.CommentList{}, errcode.ErrNotFound
		}
		com.User.ID = user.ID
		com.User.Name = user.Name
		com.User.FollowCount = user.FollowCount
		com.User.FollowerCount = user.FollowerCount
		_, err = dao.Group.Mysql.GetUserFollower(ctx, db.GetUserFollowerParams{
			FromID: userID,
			ToID:   v.UserID,
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				com.User.IsFollow = false
			}
			global.Logger.Error(err.Error())
			return reply.CommentList{}, errcode.ErrServer
		} else {
			com.User.IsFollow = true
		}
		result.CommentList = append(result.CommentList, com)
	}

	return result, nil
}
