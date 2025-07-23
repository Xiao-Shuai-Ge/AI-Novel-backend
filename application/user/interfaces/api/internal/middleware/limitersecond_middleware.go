package middleware

import (
	"Ai-Novel/common/middleware"
	"net/http"
)

type LimiterSecondMiddleware struct {
}

func NewLimiterSecondMiddleware() *LimiterSecondMiddleware {
	return &LimiterSecondMiddleware{}
}

func (m *LimiterSecondMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.LimiterMiddlewareEverySecond(next)
}
