package logic

import (
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropAlgoTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropAlgoTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropAlgoTaskLogic {
	return &DropAlgoTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropAlgoTaskLogic) DropAlgoTask(req *types.DropAlgoTaskRequest) (resp *types.BaseReply, err error) {
	ctx := l.ctx
	task, err := l.svcCtx.MysqlClient.GetAlgoTaskByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.AlgoService.StopAlgo(task.AppID)
	if err != nil {
		logx.Errorf("fail to stop algo, err: %v", err)
		return nil, err
	}
	_, _ = l.svcCtx.MysqlClient.DropAlgoTaskByID(l.ctx, req.ID)
	return &types.BaseReply{}, nil
}
