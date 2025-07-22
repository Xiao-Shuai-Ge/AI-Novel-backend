package middleware

import (
	"Ai-Novel/common/middleware"
	"net/http"
)

type LimiterMiddleware struct {
}

func NewLimiterMiddleware() *LimiterMiddleware {
	return &LimiterMiddleware{}
}

func (m *LimiterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.LimiterMiddlewareEveryMinute(next)
}
