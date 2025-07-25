package character

import (
	"Ai-Novel/common/jwtx"
	"context"
	"errors"
	"strconv"

	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCharacterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加角色
func NewCreateCharacterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCharacterLogic {
	return &CreateCharacterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCharacterLogic) CreateCharacter(req *types.CreateCharacterReq) (resp *types.CreateCharacterResp, err error) {
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
	// a. 标题和头像
	if len(req.Name) > 63 || len(req.Name) <= 0 {
		err = errors.New("角色名称长度必须大于0且小于等于63")
		return
	}
	if len(req.Avatar) > 255 {
		err = errors.New("角色头像长度必须大于0且小于等于255")
		return
	}
	// b. 总结不能超过 2000 字
	if len(req.Summary) > 2000 {
		err = errors.New("总结不能超过 2000 字")
		return
	}
	// 调用 Novel 服务
	id, err := l.svcCtx.NovelApp.CreateCharacter(l.ctx, novelID, userID, req.Name, req.Avatar, req.Summary)
	if err != nil {
		return nil, err
	}
	resp = &types.CreateCharacterResp{
		ID: id,
	}
	return
}
