package login

import (
	"context"

	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新 token
func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {
	// 调用登录服务
	atoken, err := l.svcCtx.LoginApp.RefreshToken(l.ctx, req.Rtoken)
	if err != nil {
		return nil, err
	}
	resp = &types.RefreshTokenResp{
		Atoken: atoken,
	}
	return
}
