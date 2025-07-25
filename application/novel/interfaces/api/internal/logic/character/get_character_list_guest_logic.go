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

type GetCharacterListGuestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说角色列表(游客)
func NewGetCharacterListGuestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCharacterListGuestLogic {
	return &GetCharacterListGuestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCharacterListGuestLogic) GetCharacterListGuest(req *types.GetCharacterListReq) (resp *types.GetCharacterListResp, err error) {
	// id 转 int64
	ID, err := strconv.ParseInt(req.NovelID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	chapters, err := l.svcCtx.NovelApp.GetCharacterList(l.ctx, ID, codex.UNLOGIN_USER_ID)
	if err != nil {
		return nil, err
	}
	// 转换成 GetChapterListResp
	resp = &types.GetCharacterListResp{}
	for _, character := range chapters {
		resp.List = append(resp.List, types.CharacterLite{
			ID:      character.ID,
			NovelID: character.NovelID,
			Name:    character.Name,
		})
	}
	return
}
