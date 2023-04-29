package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAlgoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAlgoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAlgoLogic {
	return &QueryAlgoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAlgoLogic) QueryAlgo(in *algo.Empty) (*algo.AlgoReply, error) {
	// todo: add your logic here and delete this line

	return &algo.AlgoReply{}, nil
}
