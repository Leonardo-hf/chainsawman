package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AvgClusteringLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAvgClusteringLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AvgClusteringLogic {
	return &AvgClusteringLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AvgClusteringLogic) AvgClustering(in *algo.BaseReq) (*algo.MetricsReply, error) {
	// todo: add your logic here and delete this line

	return &algo.MetricsReply{}, nil
}
