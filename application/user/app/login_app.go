package app

import (
	"Ai-Novel/application/user/domain/services"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/common/codex"
	"Ai-Novel/common/email"
	"Ai-Novel/common/jwtx"
	"Ai-Novel/common/utils/snowflake"
	"context"
	"errors"
)

type LoginApp struct {
	Repo          *repo.UserRepo
	EmailSender   email.Sender
	JWT           jwtx.JWT
	SnowflakeNode *snowflake.Node
}

func NewLoginApp(repo *repo.UserRepo, emailSender email.Sender, jwt jwtx.JWT, snowflakeNode *snowflake.Node) LoginApp {
	return LoginApp{
		Repo:          repo,
		EmailSender:   emailSender,
		JWT:           jwt,
		SnowflakeNode: snowflakeNode,
	}
}

// SendCaptcha
//
//	@Description: 发送验证码
//	@receiver a
//	@param email
//	@return err
func (a *LoginApp) SendCaptcha(ctx context.Context, email string) (err error) {
	// 注册验证码服务
	s := services.NewCaptchaService(ctx, a.Repo, a.EmailSender)
	// 发送验证码
	err = s.SendCaptcha(email)
	if err != nil {
		err = errors.New("发送验证码失败")
		return
	}
	return
}

// Register
//
//	@Description: 注册
//	@receiver a
//	@param ctx
//	@param email
//	@param password
//	@param captchaCode
//	@return err
func (a *LoginApp) Register(ctx context.Context, email, password, captchaCode string) (token string, err error) {
	// 1. 注册验证码服务
	captchaService := services.NewCaptchaService(ctx, a.Repo, a.EmailSender)
	// 2. 验证验证码
	yes, err := captchaService.VerifyCaptcha(email, captchaCode)
	if err != nil {
		err = errors.New("内部错误")
		return
	}
	// 3. 如果验证码不符合，返回错误
	if !yes {
		err = errors.New("验证码错误")
		return
	}
	// 4. 注册登录服务
	loginService := services.NewLoginService(ctx, a.Repo, a.JWT, a.SnowflakeNode)
	// 5. 注册用户
	user, err := loginService.Register(email, password)
	if err != nil {
		err = errors.New("内部错误")
		return
	}
	// 6. 获得 token
	token, err = loginService.GetAToken(user.ID)
	if err != nil {
		err = errors.New("生成token错误")
		return
	}
	// 7. 返回
	return
}

// Login
//
//	@Description: 登录
//	@receiver a
//	@param ctx
//	@param email
//	@param password
//	@param isRemember
//	@return atoken
//	@return rtoken
//	@return err
func (a *LoginApp) Login(ctx context.Context, email, password string, isRemember bool) (atoken string, rtoken string, err error) {
	// 1. 注册登录服务
	loginService := services.NewLoginService(ctx, a.Repo, a.JWT, a.SnowflakeNode)
	// 2. 验证密码
	user, err := loginService.Login(email, password)
	if errors.Is(err, codex.ACCOUNT_OR_PASSWORD_ERROR) {
		err = errors.New("密码或密码错误")
		return
	} else if err != nil {
		err = errors.New("内部错误")
		return
	}
	// 3. 生成 atoken
	atoken, err = loginService.GetAToken(user.ID)
	if err != nil {
		err = errors.New("生成atoken错误")
		return
	}
	// 4. 生成 rtoken
	if isRemember {
		rtoken, err = loginService.GetRToken(user.ID)
		if err != nil {
			err = errors.New("生成rtoken错误")
			return
		}
	}
	// 5. 返回
	return
}

// RefreshToken
//
//	@Description: 刷新token
//	@receiver a
//	@param ctx
//	@param rtoken
//	@return atoken
//	@return err
func (a *LoginApp) RefreshToken(ctx context.Context, rtoken string) (atoken string, err error) {
	// 1. 注册登录服务
	loginService := services.NewLoginService(ctx, a.Repo, a.JWT, a.SnowflakeNode)
	// 2. 解析 rtoken
	userID, err := loginService.ParseRToken(rtoken)
	if errors.Is(err, codex.RTOKEN_EXPIRED) {
		return
	} else if err != nil {
		err = errors.New("验证rtoken错误")
		return
	}
	// 3. 生成 atoken
	atoken, err = loginService.GetAToken(userID)
	if err != nil {
		err = errors.New("生成atoken错误")
		return
	}
	// 4. 返回
	return
}
