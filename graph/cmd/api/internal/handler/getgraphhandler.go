package handler

import (
	"net/http"

	"chainsawman/graph/api/internal/logic"
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getGraphHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetGraphLogic(r.Context(), svcCtx)
		resp, err := l.GetGraph(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
