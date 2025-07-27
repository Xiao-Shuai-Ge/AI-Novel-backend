package ai

import (
	"net/http"

	"Ai-Novel/application/ai/interfaces/api/internal/logic/ai"
	"Ai-Novel/application/ai/interfaces/api/internal/svc"
	"Ai-Novel/application/ai/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 生成小说篇章总结
func ChapterSummaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChapterSummaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ai.NewChapterSummaryLogic(r.Context(), svcCtx)
		resp, err := l.ChapterSummary(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
