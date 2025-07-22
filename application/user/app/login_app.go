package app

import (
	"Ai-Novel/application/user/domain/services"
	"Ai-Novel/common/email"
	"Ai-Novel/common/jwtx"
	"Ai-Novel/common/utils/snowflake"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LoginApp struct {
	DB            *gorm.DB
	Rdb           *redis.Client
	EmailSender   email.Sender
	JWT           jwtx.JWT
	SnowflakeNode *snowflake.Node
}

func NewLoginApp(db *gorm.DB, rdb *redis.Client, emailSender email.Sender, jwt jwtx.JWT, snowflakeNode *snowflake.Node) LoginApp {
	return LoginApp{
		DB:            db,
		Rdb:           rdb,
		EmailSender:   emailSender,
		JWT:           jwt,
		SnowflakeNode: snowflakeNode,
	}
}

// SendCaptcha
//
//	@Description: 发送验证码
//	@receiver app
//	@param email
//	@return err
func (app *LoginApp) SendCaptcha(ctx context.Context, email string) (err error) {
	// 注册验证码服务
	s := services.NewCaptchaService(ctx, app.Rdb, app.EmailSender)
	// 发送验证码
	err = s.SendCaptcha(email)
	return
}

// Register
//
//	@Description: 注册
//	@receiver app
//	@param ctx
//	@param email
//	@param password
//	@param captchaCode
//	@return err
func (app *LoginApp) Register(ctx context.Context, email, password, captchaCode string) (token string, err error) {
	// 1. 注册验证码服务
	captchaService := services.NewCaptchaService(ctx, app.Rdb, app.EmailSender)
	// 2. 验证验证码
	yes, err := captchaService.VerifyCaptcha(email, captchaCode)
	if err != nil {
		return
	}
	// 3. 如果验证码不符合，返回错误
	if !yes {
		err = errors.New("验证码错误")
		return
	}
	// 4. 注册登录服务
	loginService := services.NewLoginService(ctx, app.DB, app.Rdb, app.JWT, app.SnowflakeNode)
	// 5. 注册用户
	err = loginService.Register(email, password)
	if err != nil {
		return
	}
	// 6. 获得 token
	token, err = loginService.GetATokenByEmail(email)
	if err != nil {
		return
	}
	// 7. 返回
	return
}
