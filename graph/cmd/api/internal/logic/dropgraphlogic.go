package logic

import (
	"context"

	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropGraphLogic {
	return &DropGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropGraphLogic) DropGraph(req *types.DropRequest) (resp *types.BaseReply, err error) {
	err = l.svcCtx.NebulaClient.DropGraph(req.Graph)
	return nil, err
}
