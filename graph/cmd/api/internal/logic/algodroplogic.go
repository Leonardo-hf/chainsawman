package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlgoDropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlgoDropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlgoDropLogic {
	return &AlgoDropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlgoDropLogic) AlgoDrop(req *types.AlgoIDRequest) (resp *types.BaseReply, err error) {
	_, err = l.svcCtx.MysqlClient.DropAlgoByID(l.ctx, req.AlgoID)
	return nil, err
}
