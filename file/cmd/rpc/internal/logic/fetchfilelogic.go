package logic

import (
	"chainsawman/file/cmd/rpc/internal/svc"
	"chainsawman/file/cmd/rpc/types/rpc/file"
	"io"

	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type FetchFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFetchFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchFileLogic {
	return &FetchFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FetchFileLogic) FetchFile(in *file.IDReq) (*file.FileReply, error) {
	r, err := l.svcCtx.OSSClient.FetchAlgo(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &file.FileReply{
		Name: in.Id,
		Size: int64(len(bytes)),
		Data: bytes,
	}, nil
}
