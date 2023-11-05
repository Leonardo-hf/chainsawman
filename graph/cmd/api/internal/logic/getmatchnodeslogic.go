package logic

import (
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMatchNodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMatchNodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMatchNodesLogic {
	return &GetMatchNodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMatchNodesLogic) GetMatchNodes(req *types.GetMatchNodesRequest) (resp *types.GetMatchNodesReply, err error) {
	group, err := l.svcCtx.MysqlClient.GetGroupByGraphId(l.ctx, req.GraphID)
	if err != nil {
		return nil, err
	}
	nodes, err := l.svcCtx.NebulaClient.MatchNodes(req.GraphID, req.Keywords, group)
	if err != nil {
		return nil, err
	}
	resp = &types.GetMatchNodesReply{
		MatchNodePacks: make([]*types.MatchNodePacks, 0),
	}
	for tag, v := range nodes {
		// 复制一份
		match := make([]*types.MatchNode, len(v))
		for i, vi := range v {
			match[i] = &types.MatchNode{
				Id:          vi.Id,
				PrimaryAttr: vi.PrimaryAttr,
			}
		}
		resp.MatchNodePacks = append(resp.MatchNodePacks, &types.MatchNodePacks{
			Tag:   tag,
			Match: match,
		})
	}
	return resp, nil
}
