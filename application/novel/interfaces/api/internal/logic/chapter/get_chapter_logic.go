package chapter

import (
	"Ai-Novel/common/jwtx"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取篇章
func NewGetChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChapterLogic {
	return &GetChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChapterLogic) GetChapter(req *types.GetChapterReq) (resp *types.GetChapterResp, err error) {
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
	chapter, err := l.svcCtx.NovelApp.GetChapter(l.ctx, ID, userID)
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
