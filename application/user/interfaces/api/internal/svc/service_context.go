package svc

import (
	"Ai-Novel/application/user/app"
	"Ai-Novel/application/user/interfaces/api/internal/config"
	"Ai-Novel/application/user/interfaces/api/internal/middleware"
	"Ai-Novel/common/email"
	"Ai-Novel/common/gormx"
	"Ai-Novel/common/redisx"
	"Ai-Novel/common/zlog/dbLogger"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	CorsMiddleware rest.Middleware
	EmailSender    email.Sender
	LoginApp       app.LoginApp
	DB             *gorm.DB
	Rdb            *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库
	db := gormx.MustOpen(c.Mysql, dbLogger.New())
	rdb := redisx.MustOpen(c.Redis)

	// 初始化邮件发送器
	emailSender := email.NewEmailSender(c.EmailConf.Host, c.EmailConf.Port, c.EmailConf.Username, c.EmailConf.Password)

	return &ServiceContext{
		Config:         c,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
		EmailSender:    emailSender,
		LoginApp:       app.NewLoginApp(rdb, emailSender),
		DB:             db,
		Rdb:            rdb,
	}
}
