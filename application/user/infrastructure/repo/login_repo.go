package repo

import (
	"Ai-Novel/application/user/domain/entity"
	"Ai-Novel/common/model"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	CAPCHA_EXPIRE_TIME = 10 * time.Minute

	REDIS_CAPTCHA_KEY = "login:captcha:%s"
)

func (r *UserRepo) Register(user entity.User) error {
	// 转换数据库模型
	DbUser := user.Transform()
	// 保存到数据库
	err := r.db.Create(&DbUser).Error
	return err
}

func (r *UserRepo) GetUser(id int64) (User entity.User, err error) {
	// 1. 从数据库中查询用户
	var DbUser model.User
	err = r.db.Where("id =?", id).First(&DbUser).Error
	if err != nil {
		return
	}
	return entity.Form(&DbUser), nil
}

func (r *UserRepo) GetUserByEmail(email string) (User entity.User, err error) {
	// 1. 从数据库中查询用户
	var DbUser model.User
	err = r.db.Where("email =?", email).First(&DbUser).Error
	if err != nil {
		return
	}
	return entity.Form(&DbUser), nil
}

func (r *UserRepo) SaveCaptcha(ctx context.Context, captcha entity.Captcha) error {
	// 1. 存储到 redis 中
	err := r.rdb.Set(ctx, fmt.Sprintf(REDIS_CAPTCHA_KEY, captcha.Email), captcha.Code, CAPCHA_EXPIRE_TIME).Err()
	return err
}

func (r *UserRepo) GetCaptcha(ctx context.Context, email string) (captcha entity.Captcha, err error) {
	// 1. 从 redis 读取验证码
	var value string
	value, err = r.rdb.Get(ctx, fmt.Sprintf(REDIS_CAPTCHA_KEY, email)).Result()
	if errors.Is(err, redis.Nil) {
		// 如果验证码不存在，则返回错误
		err = errors.New("验证码不存在")
		return
	}
	captcha = entity.NewCaptcha(email)
	captcha.SetCode(value)
	return
}

func (r *UserRepo) DeleteCaptcha(ctx context.Context, email string) error {
	// 1. 从 redis 中删除验证码
	return r.rdb.Del(ctx, fmt.Sprintf(REDIS_CAPTCHA_KEY, email)).Err()
}
