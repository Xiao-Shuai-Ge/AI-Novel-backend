package services

import (
	"Ai-Novel/application/user/domain/entity"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/common/codex"
	"Ai-Novel/common/jwtx"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog"
	"context"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const (
	ATOKEN_EXPIRE_TIME = 2 * time.Hour
	RTOKEN_EXPIRE_TIME = 2 * 7 * 24 * time.Hour
)

type LoginService struct {
	ctx           context.Context
	Repo          *repo.UserRepo
	JWT           jwtx.JWT
	SnowFlakeNode *snowflake.Node
}

func NewLoginService(ctx context.Context, repo *repo.UserRepo, JWT jwtx.JWT, SnowFlakeNode *snowflake.Node) LoginService {
	return LoginService{
		ctx:           ctx,
		Repo:          repo,
		JWT:           JWT,
		SnowFlakeNode: SnowFlakeNode,
	}
}

func (s LoginService) Register(email string, password string) (user entity.User, err error) {
	// 1. 创建 user 实体并初始化信息
	user = entity.NewUser(email, password)
	// a. 加密密码
	err = user.EncryptPassword()
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "加密密码失败", err)
		return
	}
	// b. 设置唯一ID
	id := s.SnowFlakeNode.Generate()
	user.SetID(int64(id))
	// c. 设置默认用户名
	user.DefaultUsername()
	// 2. 创建账号
	err = s.Repo.Register(user)
	if err != nil {
		return
	}
	return
}

func (s LoginService) Login(email string, password string) (user entity.User, err error) {
	// 1. 用 email 查找用户
	user, err = s.Repo.GetUserByEmail(email)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "查找用户失败", err)
		err = codex.ACCOUNT_OR_PASSWORD_ERROR
		return
	}
	// 2. 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "验证密码失败", err)
		err = codex.ACCOUNT_OR_PASSWORD_ERROR
		return
	}
	return
}

func (s LoginService) ParseRToken(rtoken string) (id int64, err error) {
	// 1. 解析 rtoken
	data, err := s.JWT.IdentifyToken(rtoken)
	if err != nil {
		// 如果已过期也会报错，在外层判断
		return
	}
	// 2. data.Userid 转成 int64
	id, err = strconv.ParseInt(data.Userid, 10, 64)
	if err != nil {
		return
	}
	// 3. 返回
	return
}

func (s LoginService) GetAToken(id int64) (token string, err error) {
	// 1. 用 id 查找用户
	user, err := s.Repo.GetUser(id)

	// 2. 生成 atoken
	atoken, err := s.JWT.GenAtoken(strconv.FormatInt(user.ID, 10), ATOKEN_EXPIRE_TIME)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "生成 atoken 失败", err)
		return
	}
	token = atoken
	return
}

func (s LoginService) GetRToken(id int64) (token string, err error) {
	// 1. 用 id 查找用户
	user, err := s.Repo.GetUser(id)

	// 2. 生成 atoken
	rtoken, err := s.JWT.GenRtoken(strconv.FormatInt(user.ID, 10), RTOKEN_EXPIRE_TIME)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "生成 rtoken 失败", err)
		return
	}
	token = rtoken
	return
}
