package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClosenessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClosenessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClosenessLogic {
	return &ClosenessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClosenessLogic) Closeness(in *algo.BaseReq) (*algo.RankReply, error) {
	// todo: add your logic here and delete this line

	return &algo.RankReply{}, nil
}
