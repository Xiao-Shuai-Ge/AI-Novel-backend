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

type GetCharacterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说角色列表
func NewGetCharacterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCharacterListLogic {
	return &GetCharacterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCharacterListLogic) GetCharacterList(req *types.GetCharacterListReq) (resp *types.GetCharacterListResp, err error) {
	// 获取 操作者ID
	userIDStr := l.ctx.Value(jwtx.USER_ID_KEY).(string)
	// 转换成 int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// id 转 int64
	ID, err := strconv.ParseInt(req.NovelID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	chapters, err := l.svcCtx.NovelApp.GetCharacterList(l.ctx, ID, userID)
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
