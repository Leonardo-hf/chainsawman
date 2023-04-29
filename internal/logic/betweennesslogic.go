package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BetweennessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBetweennessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BetweennessLogic {
	return &BetweennessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BetweennessLogic) Betweenness(in *algo.BaseReq) (*algo.RankReply, error) {
	// todo: add your logic here and delete this line

	return &algo.RankReply{}, nil
}
