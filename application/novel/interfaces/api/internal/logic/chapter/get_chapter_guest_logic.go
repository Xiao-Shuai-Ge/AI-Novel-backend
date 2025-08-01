package chapter

import (
	"Ai-Novel/common/codex"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChapterGuestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取篇章(游客)
func NewGetChapterGuestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChapterGuestLogic {
	return &GetChapterGuestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChapterGuestLogic) GetChapterGuest(req *types.GetChapterReq) (resp *types.GetChapterResp, err error) {
	// id 转 int64
	ID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	chapter, err := l.svcCtx.NovelApp.GetChapter(l.ctx, ID, codex.UNLOGIN_USER_ID)
	if err != nil {
		return nil, err
	}
	resp = &types.GetChapterResp{
		Chapter: types.Chapter{
			ID:      chapter.ID,
			NovelID: chapter.NovelID,
			Title:   chapter.Title,
			Content: chapter.Content,
			Summary: chapter.Summary,
		},
	}
	return
}
