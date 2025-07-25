package middleware

import "net/http"

type LimiterSecond10Middleware struct {
}

func NewLimiterSecond10Middleware() *LimiterSecond10Middleware {
	return &LimiterSecond10Middleware{}
}

func (m *LimiterSecond10Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
