package handler

import (
	"net/http"

	"chainsawman/graph/api/internal/logic"
	"chainsawman/graph/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getAllGraphHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetAllGraphLogic(r.Context(), svcCtx)
		resp, err := l.GetAllGraph()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
