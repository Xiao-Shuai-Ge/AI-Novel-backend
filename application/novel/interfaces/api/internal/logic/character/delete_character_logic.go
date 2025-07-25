package character

import (
	"Ai-Novel/common/jwtx"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCharacterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除角色
func NewDeleteCharacterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCharacterLogic {
	return &DeleteCharacterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCharacterLogic) DeleteCharacter(req *types.DeleteCharacterReq) (resp *types.DeleteCharacterResp, err error) {
	// 获取 操作者ID
	userIDStr := l.ctx.Value(jwtx.USER_ID_KEY).(string)
	// 转换成 int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	err = l.svcCtx.NovelApp.DeleteCharacter(l.ctx, id, userID)
	if err != nil {
		return nil, err
	}
	resp = &types.DeleteCharacterResp{
		Msg: "success",
	}
	return
}
