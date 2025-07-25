package base

import (
	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 测试接口(ping)
func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.PingReq) (resp *types.PingResp, err error) {
	resp = &types.PingResp{
		Msg: "pong",
	}

	return
}
