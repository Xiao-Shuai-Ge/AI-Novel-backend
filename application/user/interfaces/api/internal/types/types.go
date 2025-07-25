// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package types

type CaptchaReq struct {
	Email string `json:"email"` // 接受验证码的邮箱
}

type CaptchaResp struct {
	Msg string `json:"msg"` // 响应信息
}

type GetUserReq struct {
	Id string `form:"id"` // 用户 id
}

type GetUserResp struct {
	Id       int64  `json:"id,string"` // 用户 id
	Username string `json:"username"`  // 用户昵称
	Avatar   string `json:"avatar"`    // 用户头像
}

type GetUserSelfReq struct {
}

type GetUserSelfResp struct {
	Id       int64  `json:"id,string"` // 用户 id
	Username string `json:"username"`  // 用户昵称
	Avatar   string `json:"avatar"`    // 用户头像
}

type LoginReq struct {
	Email      string `json:"email"`       // 登录邮箱
	Password   string `json:"password"`    // 登录密码
	IsRemember bool   `json:"is_remember"` // 是否记住登录状态
}

type LoginResp struct {
	Atoken string `json:"atoken"` // 登录 token
	Rtoken string `json:"rtoken"` // 刷新 token
}

type ModifyUserReq struct {
	Id       string `json:"id"`       // 用户 id
	Username string `json:"username"` // 用户昵称
	Avatar   string `json:"avatar"`   // 用户头像
}

type ModifyUserResp struct {
	Msg string `json:"msg"` // 响应信息
}

type PingJWTReq struct {
}

type PingJWTResp struct {
	Msg string `json:"msg"` // 响应信息
}

type PingReq struct {
}

type PingResp struct {
	Msg string `json:"msg"` // 响应信息
}

type RefreshTokenReq struct {
	Rtoken string `json:"rtoken"` // 刷新 token
}

type RefreshTokenResp struct {
	Atoken string `json:"atoken"` // 登录 token
}

type RegisterReq struct {
	Email    string `json:"email"`    // 注册邮箱
	Password string `json:"password"` // 注册密码
	Captcha  string `json:"captcha"`  // 验证码
}

type RegisterResp struct {
	Atoken string `json:"atoken"` // 登录 token
}
