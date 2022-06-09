package app

import (
	"encoding/json"
	"net/http"

	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

// State 状态码
type State struct {
	Code int    `json:"status_code"` // 状态码，0-成功，其他值-失败
	Msg  string `json:"status_msg"`  // 返回状态描述
}

type Reply struct {
	ctx *gin.Context
}

func NewReply(ctx *gin.Context) *Reply {
	return &Reply{ctx: ctx}
}

func (rep *Reply) send(err errcode.Err, data interface{}) {
	state := State{
		Code: err.ECode(),
		Msg:  err.Error(),
	}
	var result = map[string]interface{}{}
	if data != nil {
		js2, _ := json.Marshal(data)
		_ = json.Unmarshal(js2, &result)
	}
	js1, _ := json.Marshal(state)
	_ = json.Unmarshal(js1, &result)
	rep.ctx.JSON(http.StatusOK, result)
}

func (rep *Reply) SendData(data interface{}) {
	rep.send(errcode.StatusOK, data)
}

func (rep *Reply) SendErr(data interface{}, err errcode.Err) {
	rep.send(err, data)
}
