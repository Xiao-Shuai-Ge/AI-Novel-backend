package svc

import (
	"Ai-Novel/application/novel/app"
	"Ai-Novel/application/novel/infrastructure/repo"
	"Ai-Novel/application/novel/interfaces/api/internal/config"
	"Ai-Novel/application/novel/interfaces/api/internal/middleware"
	"Ai-Novel/common/email"
	"Ai-Novel/common/gormx"
	"Ai-Novel/common/jwtx"
	"Ai-Novel/common/redisx"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog/dbLogger"
	"github.com/zeromicro/go-zero/rest"
	"time"
)

type ServiceContext struct {
	Config      config.Config
	JWT         jwtx.JWT
	EmailSender email.Sender
	NovelApp    app.NovelApp

	NovelRepo *repo.NovelRepo

	SnowflakeNode *snowflake.Node

	// 中间件
	CorsMiddleware  rest.Middleware
	LimiterSecond   rest.Middleware
	LimiterSecond10 rest.Middleware
	LimiterMinute   rest.Middleware
	LimiterMinute10 rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库
	db := gormx.MustOpen(c.Mysql, dbLogger.New())
	rdb := redisx.MustOpen(c.Redis)

	novelRepo := repo.NewNovelRepo(db, rdb)

	// 初始化ID生成器
	// 目前先用时间戳作为唯一ID，后续保证集群唯一性
	id := time.Now().UnixMilli() % (1 << 10)
	snowflakeNode, err := snowflake.NewNode(id)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:        c,
		NovelApp:      app.NewNovelApp(novelRepo, snowflakeNode),
		NovelRepo:     novelRepo,
		SnowflakeNode: snowflakeNode,

		CorsMiddleware:  middleware.NewCorsMiddleware().Handle,
		LimiterSecond10: middleware.NewLimiterSecond10Middleware().Handle,
		LimiterMinute10: middleware.NewLimiterMinute10Middleware().Handle,
	}
}
