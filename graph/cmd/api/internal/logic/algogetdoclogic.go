package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlgoGetDocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlgoGetDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlgoGetDocLogic {
	return &AlgoGetDocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlgoGetDocLogic) AlgoGetDoc(req *types.AlgoIDRequest) (resp *types.GetAlgoDocReply, err error) {
	doc, err := l.svcCtx.MysqlClient.GetAlgoDoc(l.ctx, req.AlgoID)
	return &types.GetAlgoDocReply{Doc: doc}, err
}
