package logic

import (
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllGraphLogic {
	return &GetAllGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllGraphLogic) GetAllGraph() (resp *types.SearchAllGraphReply, err error) {
	fmt.Println("wdnmd")
	graphs, err := l.svcCtx.NebulaClient.GetGraphs()
	if err != nil {
		return nil, err
	}
	resp = &types.SearchAllGraphReply{}
	for _, graph := range graphs {
		resp.Graphs = append(resp.Graphs, &types.Graph{
			Name:  graph.Name,
			Nodes: graph.Nodes,
			Edges: graph.Edges,
		})
	}
	resp.Graphs = append(resp.Graphs, &types.Graph{Name: "wdnmd"})
	return resp, nil
}
