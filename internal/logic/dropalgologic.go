package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropAlgoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDropAlgoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropAlgoLogic {
	return &DropAlgoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DropAlgoLogic) DropAlgo(in *algo.DropAlgoReq) (*algo.AlgoReply, error) {
	// todo: add your logic here and delete this line

	return &algo.AlgoReply{}, nil
}
