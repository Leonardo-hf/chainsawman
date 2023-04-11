package logic

import (
	"context"

	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGraphLogic {
	return &GetGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGraphLogic) GetGraph(req *types.SearchRequest) (resp *types.SearchGraphDetailReply, err error) {
	nodes, err := l.svcCtx.NebulaClient.GetNodes(req.Graph)
	if err != nil {
		return nil, err
	}
	resp = &types.SearchGraphDetailReply{}
	for _, node := range nodes {
		resp.Nodes = append(resp.Nodes, &types.Node{
			Name: node.Name,
			Desc: node.Desc,
		})
	}
	edges, err := l.svcCtx.NebulaClient.GetEdges(req.Graph)
	if err != nil {
		return nil, err
	}
	for _, edge := range edges {
		resp.Edges = append(resp.Edges, &types.Edge{
			Source: edge.Source,
			Target: edge.Target,
		})
	}
	return resp, nil
}
