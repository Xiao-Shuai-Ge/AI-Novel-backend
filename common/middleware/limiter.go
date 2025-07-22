package middleware

import (
	"Ai-Novel/common/zlog"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"path"
	"sync"
	"time"
)

// 定义限流桶
var limiters sync.Map

// LimiterMiddlewareEverySecond10 每秒十次请求频率请求
func LimiterMiddlewareEverySecond10(next http.HandlerFunc) http.HandlerFunc {
	return LimiterMiddleware(next, rate.Every(1*time.Second)*10, 30)
}

// LimiterMiddlewareEverySecond 每秒一次请求频率请求
func LimiterMiddlewareEverySecond(next http.HandlerFunc) http.HandlerFunc {
	return LimiterMiddleware(next, rate.Every(1*time.Second), 3)
}

// LimiterMiddlewareEveryMinute10 每分钟十次请求频率请求
func LimiterMiddlewareEveryMinute10(next http.HandlerFunc) http.HandlerFunc {
	return LimiterMiddleware(next, rate.Every(1*time.Minute)*10, 30)
}

// LimiterMiddlewareEveryMinute 每分钟一次请求频率请求
func LimiterMiddlewareEveryMinute(next http.HandlerFunc) http.HandlerFunc {
	return LimiterMiddleware(next, rate.Every(1*time.Minute), 3)
}

// LimiterMiddleware 限流中间件
func LimiterMiddleware(next http.HandlerFunc, ra rate.Limit, b int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// 提取干净的API路径 (去掉查询参数)
		apiPath := cleanPath(r.URL.Path)
		// 获取真实客户端IP
		ip := realIP(r)

		// 生成唯一键 (IP + 配置)
		key := fmt.Sprintf("%s|%s", ip, apiPath)
		zlog.Debugf("生成 key 为: %s", key)

		// 获取或创建限流器
		limiter, ok := limiters.LoadOrStore(key, rate.NewLimiter(ra, b))
		if !ok {
			limiter = rate.NewLimiter(ra, b)
			limiters.Store(key, limiter)
		}

		if !limiter.(*rate.Limiter).Allow() {
			zlog.InfofCtx(ctx, "请求过于频繁! IP: %s", ip)
			writeError(w)
			return
		}

		next(w, r) // 继续处理请求
	}
}

// 清理路径（去除重复斜杠和尾随斜杠）
func cleanPath(p string) string {
	// 使用标准库规范化路径
	p = path.Clean(p)

	// 移除尾部斜杠（除非是根路径）
	if p != "/" && p[len(p)-1] == '/' {
		p = p[:len(p)-1]
	}
	return p
}

// 获取真实客户端IP (处理代理)
func realIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}
	return r.RemoteAddr
}

// 统一错误响应
func writeError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusTooManyRequests) // 429
	w.Write([]byte(fmt.Sprintf(
		`请求过于频繁!`,
	)))
}
