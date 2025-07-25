package novel

import (
	"Ai-Novel/common/codex"
	"Ai-Novel/common/jwtx"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNovelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 增加小说
func NewCreateNovelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNovelLogic {
	return &CreateNovelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNovelLogic) CreateNovel(req *types.CreateNovelReq) (resp *types.CreateNovelResp, err error) {
	// 获取 创建者ID
	userIDStr := l.ctx.Value(jwtx.USER_ID_KEY).(string)
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		err = errors.New("参数类型错误")
		return
	}
	// 数据校验
	// a. 标题和封面地址
	if len(req.Title) > 255 || len(req.Title) <= 0 {
		err = errors.New("标题长度必须大于0且小于等于255")
		return
	}
	if len(req.Avatar) > 255 || len(req.Avatar) <= 0 {
		err = errors.New("封面地址长度必须大于0且小于等于255")
		return
	}
	// b. 总结不能超过 2000 字
	if len(req.Summary) > 2000 {
		err = errors.New("总结不能超过 2000 字")
		return
	}
	// c. 状态
	if req.Status >= codex.NOVEL_STATUS_COUNT || req.Status < 0 {
		err = errors.New("状态码不存在")
		return
	}
	// 调用 Novel 服务
	id, err := l.svcCtx.NovelApp.CreateNovel(l.ctx, req.Title, req.Avatar, userID, req.Summary, req.Status, req.IsPublic)
	if err != nil {
		return nil, err
	}
	resp = &types.CreateNovelResp{
		ID: id,
	}
	return
}
