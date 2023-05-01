package logic

import (
	"chainsawman/graph/model"
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmptyGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmptyGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmptyGraphLogic {
	return &CreateEmptyGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmptyGraphLogic) CreateEmptyGraph(req *types.UploadEmptyRequest) (resp *types.SearchGraphReply, err error) {
	graph := &model.Graph{
		Name: req.Graph,
		Desc: req.Desc,
	}
	err = l.svcCtx.MysqlClient.InsertGraph(l.ctx, graph)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.NebulaClient.CreateGraph(graph.ID)
	if err != nil {
		return nil, err
	}
	return &types.SearchGraphReply{
		Graph: &types.Graph{
			Name: req.Graph,
			Desc: req.Desc,
			Id:   graph.ID,
		},
	}, nil
}
