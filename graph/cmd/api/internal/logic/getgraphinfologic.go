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

func (l *GetGraphInfoLogic) GetGraphInfo(req *types.GetGraphInfoBody) (resp *types.GetGraphInfoBody, err error) {
	// todo: add your logic here and delete this line
	name := req.Name
	graph, err := l.svcCtx.MysqlClient.GetGraphByName(l.ctx, name)
	if err != nil {
		return nil, err
	}
	return &types.GetGraphInfoBody{Id: graph.ID, Name: name}, nil
}
