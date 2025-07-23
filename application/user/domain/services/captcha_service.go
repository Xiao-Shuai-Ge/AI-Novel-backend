package services

import (
	"Ai-Novel/application/user/domain/entity"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/common/email"
	"Ai-Novel/common/zlog"
	"context"
)

type CaptchaService struct {
	ctx         context.Context
	Repo        *repo.UserRepo
	emailSender email.Sender
}

func NewCaptchaService(ctx context.Context, repo *repo.UserRepo, emailSender email.Sender) CaptchaService {
	return CaptchaService{
		ctx:         ctx,
		Repo:        repo,
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
	err = s.Repo.SaveCaptcha(s.ctx, captcha)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "保存验证码到redis失败： %s", err.Error())
	}
	return
}

func (s CaptchaService) VerifyCaptcha(email, code string) (yes bool, err error) {
	// 1. 从redis中获取验证码
	captcha, err := s.Repo.GetCaptcha(s.ctx, email)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "获取验证码失败： %s", err.Error())
		return
	}
	// 2. 验证验证码
	yes = captcha.Verify(code)
	// 3. 删除验证码
	if yes {
		err = s.Repo.DeleteCaptcha(s.ctx, email)
	}
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "删除验证码失败： %s", err.Error())
		return
	}
	// 4. 返回结果
	return
}
