package logic

import (
	"Ai-Novel/common/call/user"
	"context"

	"Ai-Novel/application/user/interfaces/rpc/internal/svc"
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

func (l *PingLogic) Ping(in *user.Ping) (*user.Pong, error) {
	return &user.Pong{
		Message: "pong," + in.Message,
	}, nil
}
