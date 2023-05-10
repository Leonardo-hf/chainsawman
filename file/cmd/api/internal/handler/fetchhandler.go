package handler

import (
	"net/http"

	"chainsawman/file/cmd/api/internal/logic"
	"chainsawman/file/cmd/api/internal/svc"
	"chainsawman/file/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func fetchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FetchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFetchLogic(r.Context(), svcCtx)
		err := l.Fetch(&req, w)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
