package logic

import (
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"

	"context"

	set "github.com/deckarep/golang-set"
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

	nodes, err := l.svcCtx.NebulaClient.GetNodes(req.Graph, req.Min)

	if err != nil {
		return nil, err
	}
	resp = &types.SearchGraphDetailReply{}
	nodeSet := set.NewThreadUnsafeSet()
	for _, node := range nodes {
		resp.Nodes = append(resp.Nodes, &types.Node{
			Name: node.Name,
			Desc: node.Desc,
		})
		nodeSet.Add(node.Name)
	}
	edges, err := l.svcCtx.NebulaClient.GetEdges(req.Graph)
	if err != nil {
		return nil, err
	}
	for _, edge := range edges {
		if nodeSet.Contains(edge.Source) || nodeSet.Contains(edge.Target) {
			resp.Edges = append(resp.Edges, &types.Edge{
				Source: edge.Source,
				Target: edge.Target,
			})
		}
	}
	return resp, nil
}
