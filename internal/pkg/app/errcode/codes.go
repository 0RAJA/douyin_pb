package errcode

// 通用错误码

var (
	StatusOK                     = NewErr(0, "成功")
	ErrParamsNotValid            = NewErr(1001, "参数绑定有误")
	ErrNotFound                  = NewErr(1002, "未找到资源")
	ErrUnauthorizedTokenTimeout  = NewErr(1003, "鉴权失败，Token 超时")
	ErrServer                    = NewErr(1004, "系统错误")
	ErrTooManyRequests           = NewErr(1005, "请求过多")
	ErrUnauthorizedAuthNotExist  = NewErr(1006, "鉴权失败, 无法解析")
	ErrUnauthorizedToken         = NewErr(1007, "鉴权失败，Token 错误")
	ErrUnauthorizedTokenGenerate = NewErr(1008, "鉴权失败，Token 生成失败")
	ErrInsufficientPermissions   = NewErr(1009, "鉴权失败,权限不足")
)

// 项目错误码

var (
	ErrUserHasExist         = NewErr(2001, "用户已存在")
	ErrUserNotExist         = NewErr(2002, "用户不存在")
	ErrNamePasswordNotMatch = NewErr(2003, "用户名或密码不正确")
	ErrLength               = NewErr(2004, "非法长度")
)
