package user

import (
	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"
	"Ai-Novel/common/jwtx"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSelfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取个人信息
func NewGetUserSelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSelfLogic {
	return &GetUserSelfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSelfLogic) GetUserSelf(req *types.GetUserSelfReq) (resp *types.GetUserSelfResp, err error) {
	// 从 ctx 拿取用户id
	userId := l.ctx.Value(jwtx.USER_ID_KEY)
	// 调用用户服务
	user, err := l.svcCtx.UserApp.GetUser(l.ctx, userId.(string))
	if err != nil {
		return
	}
	resp = &types.GetUserSelfResp{
		Id:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
	return
}
