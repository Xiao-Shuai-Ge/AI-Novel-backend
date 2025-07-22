package middleware

import (
	"Ai-Novel/common/middleware"
	"net/http"
)

type LimiterMinute10Middleware struct {
}

func NewLimiterMinute10Middleware() *LimiterMinute10Middleware {
	return &LimiterMinute10Middleware{}
}

func (m *LimiterMinute10Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.LimiterMiddlewareEveryMinute10(next)
}
