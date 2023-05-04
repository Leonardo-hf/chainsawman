package logic

import (
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"context"

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
	graphs, err := l.svcCtx.MysqlClient.GetAllGraph(l.ctx)
	if err != nil {
		return nil, err
	}
	resp = &types.SearchAllGraphReply{}
	for _, graph := range graphs {
		resp.Graphs = append(resp.Graphs, &types.Graph{
			Id:     graph.ID,
			Name:   graph.Name,
			Desc:   graph.Desc,
			Nodes:  graph.Nodes,
			Edges:  graph.Edges,
			Status: int(graph.Status),
		})
	}
	return resp, nil
}
