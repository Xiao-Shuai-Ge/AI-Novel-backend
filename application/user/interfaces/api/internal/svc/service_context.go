package svc

import (
	"Ai-Novel/application/user/app"
	"Ai-Novel/application/user/interfaces/api/internal/config"
	"Ai-Novel/application/user/interfaces/api/internal/middleware"
	"Ai-Novel/common/email"
	"Ai-Novel/common/gormx"
	"Ai-Novel/common/jwtx"
	"Ai-Novel/common/redisx"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog/dbLogger"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"time"
)

type ServiceContext struct {
	Config      config.Config
	JWT         jwtx.JWT
	EmailSender email.Sender
	LoginApp    app.LoginApp
	DB          *gorm.DB
	Rdb         *redis.Client

	SnowflakeNode *snowflake.Node

	CorsMiddleware  rest.Middleware
	LimiterMinute   rest.Middleware
	LimiterMinute10 rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库
	db := gormx.MustOpen(c.Mysql, dbLogger.New())
	rdb := redisx.MustOpen(c.Redis)

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
		LoginApp:      app.NewLoginApp(db, rdb, emailSender, jwt, snowflakeNode),
		DB:            db,
		Rdb:           rdb,
		SnowflakeNode: snowflakeNode,

		CorsMiddleware:  middleware.NewCorsMiddleware().Handle,
		LimiterMinute:   middleware.NewLimiterMinuteMiddleware().Handle,
		LimiterMinute10: middleware.NewLimiterMinute10Middleware().Handle,
	}
}
