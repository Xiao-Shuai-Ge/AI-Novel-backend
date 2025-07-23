package codex

import "errors"

// 用于存储一些公共的错误信息，方便外层判断错误类型

var (
	INTERNAL_ERROR = errors.New("内部错误")

	// 登录相关
	ACCOUNT_OR_PASSWORD_ERROR = errors.New("账号或密码错误")
	RTOKEN_EXPIRED            = errors.New("refresh token 已过期")
)
