package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PagerankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPagerankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PagerankLogic {
	return &PagerankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PagerankLogic) Pagerank(in *algo.PageRankReq) (*algo.RankReply, error) {
	// todo: add your logic here and delete this line

	return &algo.RankReply{}, nil
}
