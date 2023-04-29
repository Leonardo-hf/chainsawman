package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DegreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDegreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DegreeLogic {
	return &DegreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DegreeLogic) Degree(in *algo.BaseReq) (*algo.RankReply, error) {
	// todo: add your logic here and delete this line

	return &algo.RankReply{}, nil
}
