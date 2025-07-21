package app

import (
	"Ai-Novel/application/user/domain/services"
	"Ai-Novel/common/email"
	"context"
	"github.com/redis/go-redis/v9"
)

type LoginApp struct {
	Rdb         *redis.Client
	EmailSender email.Sender
}

func NewLoginApp(rdb *redis.Client, emailSender email.Sender) LoginApp {
	return LoginApp{
		Rdb:         rdb,
		EmailSender: emailSender,
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
