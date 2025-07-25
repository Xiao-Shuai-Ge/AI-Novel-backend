package character

import (
	"net/http"

	"Ai-Novel/application/novel/interfaces/api/internal/logic/character"
	"Ai-Novel/application/novel/interfaces/api/internal/svc"
	"Ai-Novel/application/novel/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取小说角色列表
func GetCharacterListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCharacterListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := character.NewGetCharacterListLogic(r.Context(), svcCtx)
		resp, err := l.GetCharacterList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
