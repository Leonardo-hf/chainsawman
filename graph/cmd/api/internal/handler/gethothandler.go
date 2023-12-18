package handler

import (
	"net/http"

	"chainsawman/graph/cmd/api/internal/logic"
	"chainsawman/graph/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getHotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetHotLogic(r.Context(), svcCtx)
		resp, err := l.GetHot()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
