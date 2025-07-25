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

type DeleteChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除篇章
func NewDeleteChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChapterLogic {
	return &DeleteChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteChapterLogic) DeleteChapter(req *types.DeleteChapterReq) (resp *types.DeleteChapterResp, err error) {
	// 获取 操作者ID
	userIDStr := l.ctx.Value(jwtx.USER_ID_KEY).(string)
	// 转换成 int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 调用 Novel 服务
	err = l.svcCtx.NovelApp.DeleteChapter(l.ctx, id, userID)
	if err != nil {
		return nil, err
	}
	resp = &types.DeleteChapterResp{
		Msg: "success",
	}
	return
}
