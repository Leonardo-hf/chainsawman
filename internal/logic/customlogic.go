package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomLogic {
	return &CustomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomLogic) Custom(in *algo.CustomAlgoReq) (*algo.CustomAlgoReply, error) {
	// todo: add your logic here and delete this line

	return &algo.CustomAlgoReply{}, nil
}
