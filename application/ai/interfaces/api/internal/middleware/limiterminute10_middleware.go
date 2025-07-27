package middleware

import "net/http"

type LimiterMinute10Middleware struct {
}

func NewLimiterMinute10Middleware() *LimiterMinute10Middleware {
	return &LimiterMinute10Middleware{}
}

func (m *LimiterMinute10Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
