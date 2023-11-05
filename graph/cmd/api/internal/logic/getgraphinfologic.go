package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGraphInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGraphInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGraphInfoLogic {
	return &GetGraphInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGraphInfoLogic) GetGraphInfo(req *types.GetGraphInfoRequest) (resp *types.GraphInfoReply, err error) {
	name := req.Name
	graph, err := l.svcCtx.MysqlClient.GetGraphByName(l.ctx, name)
	if err != nil {
		return nil, err
	}
	return &types.GraphInfoReply{
		Graph: &types.Graph{
			Id:       graph.ID,
			Status:   graph.Status,
			GroupID:  graph.GroupID,
			Name:     graph.Name,
			NumNode:  graph.NumNode,
			NumEdge:  graph.NumEdge,
			CreatAt:  graph.CreateTime.UnixMilli(),
			UpdateAt: graph.UpdateTime.UnixMilli(),
		},
	}, nil
}
