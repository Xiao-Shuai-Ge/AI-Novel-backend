package services

import (
	"Ai-Novel/application/user/domain/entity"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/common/jwtx"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
	ATOKEN_EXPIRE_TIME = 2 * time.Hour
	RTOKEN_EXPIRE_TIME = 2 * 7 * 24 * time.Hour
)

type LoginService struct {
	ctx           context.Context
	db            *gorm.DB
	rdb           *redis.Client
	JWT           jwtx.JWT
	SnowFlakeNode *snowflake.Node
}

func NewLoginService(ctx context.Context, db *gorm.DB, rdb *redis.Client, JWT jwtx.JWT, SnowFlakeNode *snowflake.Node) LoginService {
	return LoginService{
		ctx:           ctx,
		db:            db,
		rdb:           rdb,
		JWT:           JWT,
		SnowFlakeNode: SnowFlakeNode,
	}
}

func (s LoginService) Register(email string, password string) (err error) {
	// 1. 创建 user 实体并初始化信息
	user := entity.NewUser(email, password)
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
	// 2. 注册 repo 层
	r := repo.NewLoginRepo(s.ctx, s.db, s.rdb)
	// 3. 创建账号
	err = r.Register(user)
	if err != nil {
		return
	}
	return
}

func (s LoginService) GetATokenByEmail(email string) (token string, err error) {
	// 1. 用 email 查找用户
	r := repo.NewLoginRepo(s.ctx, s.db, s.rdb)
	user, err := r.GetUserByEmail(email)

	// 2. 生成 atoken
	atoken, err := s.JWT.GenAtoken(strconv.FormatInt(user.ID, 10), ATOKEN_EXPIRE_TIME)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "生成 atoken 失败", err)
		return
	}
	token = atoken
	return
}
