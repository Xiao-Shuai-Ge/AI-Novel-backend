package login

import (
	"context"

	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送注册邮箱验证码
func NewSendCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCaptchaLogic {
	return &SendCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCaptchaLogic) SendCaptcha(req *types.CaptchaReq) (resp *types.CaptchaResp, err error) {
	// 使用 login 服务中的 SendCaptcha 方法发送验证码
	err = l.svcCtx.LoginApp.SendCaptcha(l.ctx, req.Email)
	// 处理错误
	if err != nil {
		return
	}
	// 填入响应数据
	resp = &types.CaptchaResp{
		Msg: "验证码已发送至您的邮箱，请注意查收。",
	}
	return
}
