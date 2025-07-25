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

type CreateChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加篇章
func NewCreateChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChapterLogic {
	return &CreateChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateChapterLogic) CreateChapter(req *types.CreateChapterReq) (resp *types.CreateChapterResp, err error) {
	// 获取 创建者ID
	userIDStr := l.ctx.Value(jwtx.USER_ID_KEY).(string)
	// ID 转换成 int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	novelID, err := strconv.ParseInt(req.NovelID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 数据校验
	// a. 标题
	if len(req.Title) > 255 || len(req.Title) <= 0 {
		err = errors.New("标题长度必须大于0且小于等于255")
		return
	}
	// b. 总结不能超过 2000 字
	if len(req.Summary) > 2000 {
		err = errors.New("总结不能超过 2000 字")
		return
	}
	// c. 正文不能超过 50000 字
	if len(req.Content) > 50000 {
		err = errors.New("正文不能超过 50000 字")
		return
	}
	// 调用 Novel 服务
	id, err := l.svcCtx.NovelApp.CreateChapter(l.ctx, novelID, userID, req.Title, req.Content, req.Summary)
	if err != nil {
		return nil, err
	}
	resp = &types.CreateChapterResp{
		ID: id,
	}

	return
}
