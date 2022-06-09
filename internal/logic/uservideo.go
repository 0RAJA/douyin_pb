package logic

import (
	"database/sql"

	"github.com/0RAJA/douyin/internal/dao"
	"github.com/0RAJA/douyin/internal/global"
	"github.com/0RAJA/douyin/internal/model/request"
	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type uservideo struct{}

func (uservideo) Action(ctx *gin.Context, params request.FavoriteAction) errcode.Err {
	if params.ActionType == 1 {
		if err := dao.Group.Mysql.AddFavorite(ctx, params.VideoID); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				global.Logger.Error(err.Error())
				return errcode.ErrServer
			}
		} else {
			return errcode.ErrUserHasExist
		}
	} else {
		if err := dao.Group.Mysql.DeleteFavorite(ctx, params.VideoID); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				global.Logger.Error(err.Error())
				return errcode.ErrServer
			}
		} else {
			return errcode.ErrUserNotExist
		}
	}
	return nil

}
