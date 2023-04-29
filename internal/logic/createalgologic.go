package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAlgoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAlgoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAlgoLogic {
	return &CreateAlgoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAlgoLogic) CreateAlgo(in *algo.CreateAlgoReq) (*algo.AlgoReply, error) {
	// todo: add your logic here and delete this line

	return &algo.AlgoReply{}, nil
}
