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

type GetCharacterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说角色信息
func NewGetCharacterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCharacterLogic {
	return &GetCharacterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCharacterLogic) GetCharacter(req *types.GetCharacterReq) (resp *types.GetCharacterResp, err error) {
	// 获取 操作者ID
	userIDStr := l.ctx.Value(jwtx.USER_ID_KEY).(string)
	// 转换成 int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// id 转 int64
	ID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	character, err := l.svcCtx.NovelApp.GetCharacter(l.ctx, ID, userID)
	if err != nil {
		return nil, err
	}
	resp = &types.GetCharacterResp{
		Character: types.Character{
			ID:      character.ID,
			NovelID: character.NovelID,
			Name:    character.Name,
			Avatar:  character.Avatar,
			Summary: character.Summary,
		},
	}
	return
}
