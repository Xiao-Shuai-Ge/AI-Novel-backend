package base

import (
	"net/http"

	"Ai-Novel/application/user/interfaces/api/internal/logic/base"
	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/application/user/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// JWT测试接口(ping)
func PingJWTHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingJWTReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := base.NewPingJWTLogic(r.Context(), svcCtx)
		resp, err := l.PingJWT(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
