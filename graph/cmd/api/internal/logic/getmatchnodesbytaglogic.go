package logic

import (
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMatchNodesByTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMatchNodesByTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMatchNodesByTagLogic {
	return &GetMatchNodesByTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMatchNodesByTagLogic) GetMatchNodesByTag(req *types.GetMatchNodesByTagRequest) (resp *types.GetMatchNodesByTagReply, err error) {
	node, err := l.svcCtx.MysqlClient.GetNodeByID(l.ctx, req.NodeID)
	if err != nil {
		return nil, err
	}
	matchs, err := l.svcCtx.NebulaClient.MatchNodesByTag(req.GraphID, req.Keywords, node)
	if err != nil {
		return nil, err
	}
	matchNodes := make([]*types.MatchNode, len(matchs))
	for i, vi := range matchs {
		matchNodes[i] = &types.MatchNode{
			Id:          vi.Id,
			PrimaryAttr: vi.PrimaryAttr,
		}
	}
	return &types.GetMatchNodesByTagReply{
		MatchNodes: matchNodes,
	}, nil
}
