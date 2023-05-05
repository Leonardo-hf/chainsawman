package handler

import (
	"fmt"
	"net/http"

	"chainsawman/graph/cmd/api/internal/logic"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getNodesInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNodeReduceRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(req)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetNodesInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetNodesInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
