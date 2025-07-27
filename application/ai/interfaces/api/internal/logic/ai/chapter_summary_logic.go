package ai

import (
	"context"

	"Ai-Novel/application/ai/interfaces/api/internal/svc"
	"Ai-Novel/application/ai/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChapterSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成小说篇章总结
func NewChapterSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterSummaryLogic {
	return &ChapterSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChapterSummaryLogic) ChapterSummary(req *types.ChapterSummaryReq) (resp *types.ChapterSummaryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
