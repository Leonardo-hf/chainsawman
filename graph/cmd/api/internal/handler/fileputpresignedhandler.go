package handler

import (
	"net/http"

	"chainsawman/graph/cmd/api/internal/logic"
	"chainsawman/graph/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func filePutPresignedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFilePutPresignedLogic(r.Context(), svcCtx)
		resp, err := l.FilePutPresigned()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
