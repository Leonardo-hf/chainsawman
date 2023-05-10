package logic

import (
	"context"
	"io"
	"net/http"

	"chainsawman/file/cmd/api/internal/svc"
	"chainsawman/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchLogic {
	return &FetchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchLogic) Fetch(req *types.FetchReq, w http.ResponseWriter) error {
	r, err := l.svcCtx.OSSClient.FetchAlgo(l.ctx, req.ID)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, r)
	return err
}
