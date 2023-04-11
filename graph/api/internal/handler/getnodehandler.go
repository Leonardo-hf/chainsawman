package handler

import (
	"net/http"

	"chainsawman/graph/api/internal/logic"
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getNodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchNodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetNodeLogic(r.Context(), svcCtx)
		resp, err := l.GetNode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
