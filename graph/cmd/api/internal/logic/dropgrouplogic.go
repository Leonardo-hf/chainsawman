package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropGroupLogic {
	return &DropGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropGroupLogic) DropGroup(req *types.DropGroupRequest) (resp *types.BaseReply, err error) {
	// 删除组内图空间
	graphs, err := l.svcCtx.MysqlClient.GetGraphByGroupID(l.ctx, req.GroupID)
	if err != nil {
		return nil, err
	}
	for _, g := range graphs {
		err = l.svcCtx.NebulaClient.DropGraph(g.ID)
		if err != nil {
			return nil, err
		}
	}
	// 删除组
	_, err = l.svcCtx.MysqlClient.DropGroupByID(l.ctx, req.GroupID)
	return nil, err
}
