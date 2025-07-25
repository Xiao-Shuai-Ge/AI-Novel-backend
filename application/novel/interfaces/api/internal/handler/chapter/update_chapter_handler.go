package chapter

import (
	"net/http"

	"Ai-Novel/application/novel/interfaces/api/internal/logic/chapter"
	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改篇章
func UpdateChapterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateChapterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chapter.NewUpdateChapterLogic(r.Context(), svcCtx)
		resp, err := l.UpdateChapter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
