package svc

import (
	"Ai-Novel/application/user/app"
	"Ai-Novel/application/user/infrastructure/repo"
	"Ai-Novel/application/user/interfaces/api/internal/config"
	"Ai-Novel/application/user/interfaces/api/internal/middleware"
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
	LoginApp    app.LoginApp
	UserApp     app.UserApp

	UserRepo *repo.UserRepo

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

	userRepo := repo.NewUserRepo(db, rdb)

	// 初始化ID生成器
	// TODO 目前先用时间戳作为唯一ID，后续保证集群唯一性
	id := time.Now().UnixMilli() % (1 << 10)
	snowflakeNode, err := snowflake.NewNode(id)
	if err != nil {
		panic(err)
	}

	// jwt
	jwt := jwtx.NewJWT(c.JwtSecret)

	// 初始化邮件发送器
	emailSender := email.NewEmailSender(c.EmailConf.Host, c.EmailConf.Port, c.EmailConf.Username, c.EmailConf.Password)

	return &ServiceContext{
		Config:        c,
		JWT:           jwt,
		EmailSender:   emailSender,
		LoginApp:      app.NewLoginApp(userRepo, emailSender, jwt, snowflakeNode),
		UserApp:       app.NewUserApp(userRepo, snowflakeNode),
		UserRepo:      userRepo,
		SnowflakeNode: snowflakeNode,

		CorsMiddleware:  middleware.NewCorsMiddleware().Handle,
		LimiterSecond:   middleware.NewLimiterSecondMiddleware().Handle,
		LimiterSecond10: middleware.NewLimiterSecond10Middleware().Handle,
		LimiterMinute:   middleware.NewLimiterMinuteMiddleware().Handle,
		LimiterMinute10: middleware.NewLimiterMinute10Middleware().Handle,
	}
}
