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

type GetChapterListGusetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说篇章列表(游客)
func NewGetChapterListGusetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChapterListGusetLogic {
	return &GetChapterListGusetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChapterListGusetLogic) GetChapterListGuset(req *types.GetChapterListReq) (resp *types.GetChapterListResp, err error) {
	// id 转 int64
	ID, err := strconv.ParseInt(req.NovelID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	chapters, err := l.svcCtx.NovelApp.GetChapterList(l.ctx, ID, codex.UNLOGIN_USER_ID)
	if err != nil {
		return nil, err
	}
	// 转换成 GetChapterListResp
	resp = &types.GetChapterListResp{}
	for _, chapter := range chapters {
		resp.List = append(resp.List, types.ChapterLite{
			ID:      chapter.ID,
			NovelID: chapter.NovelID,
			Title:   chapter.Title,
		})
	}
	return
}
