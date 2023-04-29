package logic

import (
	"context"

	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VoterankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVoterankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VoterankLogic {
	return &VoterankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VoterankLogic) Voterank(in *algo.VoteRankReq) (*algo.RankReply, error) {
	// todo: add your logic here and delete this line

	return &algo.RankReply{}, nil
}
