package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropTaskLogic {
	return &DropTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropTaskLogic) DropTask(req *types.DropTaskRequest) (resp *types.BaseReply, err error) {
	ctx := l.ctx
	task, _ := l.svcCtx.MysqlClient.GetTaskByID(ctx, req.TaskID)
	_ = l.svcCtx.TaskMq.DelTaskMsg(l.ctx, task.Tid)
	_, _ = l.svcCtx.MysqlClient.DropTaskByID(l.ctx, req.TaskID)
	_, _ = l.svcCtx.RedisClient.DropTask(l.ctx, req.TaskID)
	return &types.BaseReply{}, nil
}
