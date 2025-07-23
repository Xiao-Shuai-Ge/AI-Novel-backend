package services

import (
	"Ai-Novel/application/user/domain/entity"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog"
	"context"
)

type UserService struct {
	ctx           context.Context
	Repo          *repo.UserRepo
	SnowFlakeNode *snowflake.Node
}

func NewUserService(ctx context.Context, repo *repo.UserRepo, SnowFlakeNode *snowflake.Node) UserService {
	return UserService{
		ctx:           ctx,
		Repo:          repo,
		SnowFlakeNode: SnowFlakeNode,
	}
}

func (s *UserService) GetUser(id int64) (user entity.User, err error) {
	// 1. 从数据库中获取用户信息
	user, err = s.Repo.GetUser(id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误：%v", err)
		return
	}
	return
}

func (s *UserService) ModifyUser(user entity.User) (err error) {
	// 1. 从数据库中修改用户信息
	err = s.Repo.ModifyUserProfile(user)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误：%v", err)
		return
	}
	return
}
