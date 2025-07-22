package services

import (
	"Ai-Novel/application/user/domain/entity"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/common/email"
	"Ai-Novel/common/zlog"
	"context"
	"github.com/redis/go-redis/v9"
)

type CaptchaService struct {
	ctx         context.Context
	rdb         *redis.Client
	emailSender email.Sender
}

func NewCaptchaService(ctx context.Context, rdb *redis.Client, emailSender email.Sender) CaptchaService {
	return CaptchaService{
		ctx:         ctx,
		rdb:         rdb,
		emailSender: emailSender,
	}
}

func (s CaptchaService) SendCaptcha(email string) (err error) {
	// 1. 创造验证码实体
	captcha := entity.NewCaptcha(email)
	// 2. 检查邮箱格式
	err = captcha.CheckEmail()
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "邮箱格式错误： %s", err.Error())
		return
	}
	// 3. 生成验证码
	captcha.SetRandCode()
	// 4. 发送验证码
	err = s.emailSender.Send([]string{email}, "【AI-Novel】验证码", "您的验证码为："+captcha.Code)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "发送邮箱失败： %s", err.Error())
		return
	}
	// 5. 保存验证码到redis
	r := repo.NewLoginRepo(s.ctx, nil, s.rdb)
	err = r.SaveCaptcha(captcha)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "保存验证码到redis失败： %s", err.Error())
	}
	return
}
