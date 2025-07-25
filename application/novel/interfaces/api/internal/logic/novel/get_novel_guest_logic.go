package novel

import (
	"Ai-Novel/common/codex"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNovelGuestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说信息(游客)
func NewGetNovelGuestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNovelGuestLogic {
	return &GetNovelGuestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNovelGuestLogic) GetNovelGuest(req *types.GetNovelReq) (resp *types.GetNovelResp, err error) {
	// id 转 int64
	ID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	novel, err := l.svcCtx.NovelApp.GetNovel(l.ctx, ID, codex.UNLOGIN_USER_ID)
	if err != nil {
		return nil, err
	}
	resp = &types.GetNovelResp{
		Novel: types.Novel{
			ID:       novel.ID,
			Title:    novel.Title,
			AuthorID: novel.AuthorID,
			Summary:  novel.Summary,
			Status:   novel.Status,
			IsPublic: novel.IsPublic,
		},
	}
	return
}
