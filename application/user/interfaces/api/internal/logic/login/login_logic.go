package login

import (
	"context"

	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 调用登录服务
	atoken, rtoken, err := l.svcCtx.LoginApp.Login(l.ctx, req.Email, req.Password, req.IsRemember)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResp{
		Atoken: atoken,
		Rtoken: rtoken,
	}
	return
}
