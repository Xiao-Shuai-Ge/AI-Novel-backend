package repo

import (
	"Ai-Novel/application/user/domain/entity"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

var (
	CAPCHA_EXPIRE_TIME = 10 * time.Minute

	REDIS_CAPTCHA_KEY = "login:captcha:%s"
)

type LoginRepo struct {
	ctx context.Context
	db  *gorm.DB
	rdb *redis.Client
}

func NewLoginRepo(ctx context.Context, db *gorm.DB, redis *redis.Client) LoginRepo {
	return LoginRepo{
		ctx: ctx,
		db:  db,
		rdb: redis,
	}
}

func (r *LoginRepo) SaveCaptcha(captcha entity.Captcha) error {
	// 1. 存储到 redis 中
	err := r.rdb.Set(r.ctx, fmt.Sprintf(REDIS_CAPTCHA_KEY, captcha.GetEmail()), captcha.GetCode(), CAPCHA_EXPIRE_TIME).Err()
	return err
}
