package logic

import (
	"context"
	"os"
	p "path"

	"chainsawman/file/rpc/internal/svc"
	"chainsawman/file/rpc/types/rpc"

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

func (l *FetchFileLogic) FetchFile(in *rpc.IDReq) (*rpc.FileReply, error) {
	name := p.Join(l.svcCtx.Config.Path, in.Id+".csv")
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return &rpc.FileReply{
		Name: in.Id,
		Size: int64(len(data) >> 10),
		Data: data,
	}, nil
}
