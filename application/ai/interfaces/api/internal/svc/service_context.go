package svc

import (
	"Ai-Novel/application/ai/interfaces/api/internal/config"
	"Ai-Novel/application/ai/interfaces/api/internal/middleware"
	"Ai-Novel/common/jwtx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	JWT    jwtx.JWT

	LimiterMinute   rest.Middleware
	LimiterMinute10 rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// jwt
	jwt := jwtx.NewJWT(c.JwtSecret)

	return &ServiceContext{
		Config: c,
		JWT:    jwt,

		LimiterMinute:   middleware.NewLimiterMinuteMiddleware().Handle,
		LimiterMinute10: middleware.NewLimiterMinute10Middleware().Handle,
	}
}
