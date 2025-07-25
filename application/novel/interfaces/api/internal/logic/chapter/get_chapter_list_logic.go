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

type GetChapterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说篇章列表
func NewGetChapterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChapterListLogic {
	return &GetChapterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChapterListLogic) GetChapterList(req *types.GetChapterListReq) (resp *types.GetChapterListResp, err error) {
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
	chapters, err := l.svcCtx.NovelApp.GetChapterList(l.ctx, ID, userID)
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
