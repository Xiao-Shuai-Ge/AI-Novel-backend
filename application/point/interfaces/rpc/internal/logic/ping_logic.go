package logic

import (
	"Ai-Novel/common/call/user"
	"context"

	"Ai-Novel/application/point/interfaces/rpc/internal/svc"
	"Ai-Novel/application/point/interfaces/rpc/point"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *point.Ping) (*point.Pong, error) {
	// 调用 user 的 ping 服务来实现，测试不同服务间的调用
	userPong, err := l.svcCtx.UserRpc.Ping(l.ctx, &user.Ping{
		Message: "I'm user ping",
	})

	return &point.Pong{
		Message: userPong.Message,
	}, err
}
