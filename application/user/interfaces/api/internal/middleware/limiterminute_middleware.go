package middleware

import (
	"Ai-Novel/common/middleware"
	"net/http"
)

type LimiterMinuteMiddleware struct {
}

func NewLimiterMinuteMiddleware() *LimiterMinuteMiddleware {
	return &LimiterMinuteMiddleware{}
}

func (m *LimiterMinuteMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.LimiterMiddlewareEveryMinute(next)
}
