package chapter

import (
	"net/http"

	"Ai-Novel/application/novel/interfaces/api/internal/logic/chapter"
	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取篇章
func GetChapterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetChapterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chapter.NewGetChapterLogic(r.Context(), svcCtx)
		resp, err := l.GetChapter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
