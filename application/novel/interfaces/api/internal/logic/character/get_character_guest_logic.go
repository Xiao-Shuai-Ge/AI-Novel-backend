package character

import (
	"Ai-Novel/common/codex"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCharacterGuestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说角色信息(游客)
func NewGetCharacterGuestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCharacterGuestLogic {
	return &GetCharacterGuestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCharacterGuestLogic) GetCharacterGuest(req *types.GetCharacterReq) (resp *types.GetCharacterResp, err error) {
	// id 转 int64
	ID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	character, err := l.svcCtx.NovelApp.GetCharacter(l.ctx, ID, codex.UNLOGIN_USER_ID)
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
