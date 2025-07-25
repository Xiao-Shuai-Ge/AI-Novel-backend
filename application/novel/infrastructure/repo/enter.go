package repo

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type NovelRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewNovelRepo(db *gorm.DB, redis *redis.Client) *NovelRepo {
	return &NovelRepo{
		db:  db,
		rdb: redis,
	}
}
