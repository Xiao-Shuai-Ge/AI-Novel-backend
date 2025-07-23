package repo

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserRepo(db *gorm.DB, redis *redis.Client) *UserRepo {
	return &UserRepo{
		db:  db,
		rdb: redis,
	}
}
