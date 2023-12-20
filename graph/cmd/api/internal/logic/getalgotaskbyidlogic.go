package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAlgoTaskByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAlgoTaskByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAlgoTaskByIDLogic {
	return &GetAlgoTaskByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAlgoTaskByIDLogic) GetAlgoTaskByID(req *types.GetAlgoTaskRequest) (resp *types.GetAlgoTaskReply, err error) {
	exec, err := l.svcCtx.MysqlClient.GetAlgoTaskByID(l.ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &types.GetAlgoTaskReply{
		Task: &types.AlgoTask{
			Id:         exec.ID,
			CreateTime: exec.CreateTime.UnixMilli(),
			UpdateTime: exec.UpdateTime.UnixMilli(),
			GraphID:    exec.GraphID,
			Req:        exec.Params,
			Status:     exec.Status,
			AlgoID:     exec.AlgoID,
			Output:     exec.Output,
		},
	}, err
}
