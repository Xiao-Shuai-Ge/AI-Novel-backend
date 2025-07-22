package middleware

import (
	"Ai-Novel/common/middleware"
	"net/http"
	"sync"
)

type LimiterMiddleware struct {
}

func NewLimiterMiddleware() *LimiterMiddleware {
	return &LimiterMiddleware{}
}

// 定义限流桶
var limiters sync.Map

func (m *LimiterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.LimiterMiddlewareEveryMinute(next)
}
