package logic

import (
	"chainsawman/graph/model"
	"context"
	"strconv"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAlgoTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAlgoTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAlgoTaskLogic {
	return &GetAlgoTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAlgoTaskLogic) GetAlgoTask(req *types.GetAlgoTasksRequest) (resp *types.GetAlgoTasksReply, err error) {
	var tasks []*model.Exec
	if req.GraphID == 0 {
		tasks, err = l.svcCtx.MysqlClient.GetAlgoTasks(l.ctx)
	} else {
		tasks, err = l.svcCtx.MysqlClient.GetAlgoTasksByGraphId(l.ctx, req.GraphID)
	}
	if err != nil {
		return nil, err
	}
	resp = &types.GetAlgoTasksReply{
		Tasks: make([]*types.AlgoTask, 0),
	}
	// reverse
	for i := range tasks {
		task := tasks[len(tasks)-1-i]
		resp.Tasks = append(resp.Tasks, &types.AlgoTask{
			Id:         strconv.FormatInt(task.ID, 10),
			CreateTime: task.CreateTime.UnixMilli(),
			UpdateTime: task.UpdateTime.UnixMilli(),
			Status:     task.Status,
			Req:        task.Params,
			AlgoID:     task.AlgoID,
			Output:     task.Output,
		})
	}
	return resp, nil
}
