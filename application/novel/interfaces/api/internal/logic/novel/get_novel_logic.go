package novel

import (
	"Ai-Novel/common/jwtx"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNovelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取小说信息
func NewGetNovelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNovelLogic {
	return &GetNovelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNovelLogic) GetNovel(req *types.GetNovelReq) (resp *types.GetNovelResp, err error) {
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
	novel, err := l.svcCtx.NovelApp.GetNovel(l.ctx, ID, userID)
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
