package base

import (
	"context"

	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingJWTLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// JWT测试接口(ping)
func NewPingJWTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingJWTLogic {
	return &PingJWTLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingJWTLogic) PingJWT(req *types.PingJWTReq) (resp *types.PingJWTResp, err error) {
	value := l.ctx.Value("userid")

	resp = &types.PingJWTResp{
		Msg: "pong, userid: " + value.(string),
	}

	return
}
