package user

import (
	"Ai-Novel/common/jwtx"
	"context"

	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户信息
func NewModifyUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserLogic {
	return &ModifyUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyUserLogic) ModifyUser(req *types.ModifyUserReq) (resp *types.ModifyUserResp, err error) {
	// 从 ctx 拿取用户id
	userId := l.ctx.Value(jwtx.USER_ID_KEY)
	// 调用用户服务
	err = l.svcCtx.UserApp.ModifyUser(l.ctx, userId.(string), req.Username, req.Avatar)
	if err != nil {
		return
	}
	resp = &types.ModifyUserResp{
		Msg: "success",
	}
	return
}
