package logic

import (
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"context"

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

func (l *DropGraphLogic) DropGraph(req *types.DropGraphRequest) (resp *types.BaseReply, err error) {
	err = l.svcCtx.NebulaClient.DropGraph(req.GraphID)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MysqlClient.DropGraphByID(l.ctx, req.GraphID)
	return nil, err
}
