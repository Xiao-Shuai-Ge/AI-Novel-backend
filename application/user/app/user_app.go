package app

import (
	"Ai-Novel/application/user/domain/entity"
	"Ai-Novel/application/user/domain/services"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog"
	"context"
	"errors"
	"strconv"
)

type UserApp struct {
	Repo          *repo.UserRepo
	SnowflakeNode *snowflake.Node
}

func NewUserApp(repo *repo.UserRepo, snowflakeNode *snowflake.Node) UserApp {
	return UserApp{
		Repo:          repo,
		SnowflakeNode: snowflakeNode,
	}
}

func (a *UserApp) GetUser(ctx context.Context, userId string) (user entity.User, err error) {
	// 1. 将 userId 转换为 int64
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		zlog.ErrorfCtx(ctx, "参数类型错误")
		err = errors.New("参数类型错误")
		return
	}
	// 2. 注册用户服务
	userService := services.NewUserService(ctx, a.Repo, a.SnowflakeNode)
	// 3. 调用服务获取用户信息
	user, err = userService.GetUser(id)
	if err != nil {
		err = errors.New("用户不存在")
	}
	return
}

func (a *UserApp) ModifyUser(ctx context.Context, userId, username, avatar string) (err error) {
	// 1. 数据校验:
	// a. 将 userId 转换为 int64
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		zlog.ErrorfCtx(ctx, "参数类型错误")
		err = errors.New("参数类型错误")
		return
	}
	// b. 校验 username
	if len(username) > 20 || len(username) == 0 {
		zlog.ErrorfCtx(ctx, "用户名长度超过限制")
		err = errors.New("用户名长度超过限制")
		return
	}
	// c. 校验 avatar
	if len(avatar) > 512 || len(avatar) == 0 {
		zlog.ErrorfCtx(ctx, "头像长度超过限制")
		err = errors.New("头像长度超过限制")
		return
	}
	// 2. 注册用户服务
	userService := services.NewUserService(ctx, a.Repo, a.SnowflakeNode)
	// 3. 设置 user
	user := entity.NewUserByProfile(id, username, avatar)
	// 4. 调用服务修改用户信息
	err = userService.ModifyUser(user)
	if err != nil {
		err = errors.New("修改失败")
	}
	return
}
