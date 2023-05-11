package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodesInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodesInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodesInfoLogic {
	return &GetNodesInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodesInfoLogic) GetNodesInfo(req *types.GetNodeReduceRequest) (resp *types.NodesInfo, err error) {
	nodeList, err := l.svcCtx.NebulaClient.GetAllNodes(req.Id)
	if err != nil {
		return nil, err
	}
	nodes := make([]types.NodeReduce, len(nodeList))
	for i, value := range nodeList {
		temp := types.NodeReduce{
			Id:   value.ID,
			Name: value.Name,
		}
		nodes[i] = temp
	}
	return &types.NodesInfo{Nodes: nodes}, nil
}
