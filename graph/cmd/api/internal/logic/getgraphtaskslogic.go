package logic

import (
	"context"
	"strconv"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGraphTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGraphTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGraphTasksLogic {
	return &GetGraphTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGraphTasksLogic) GetGraphTasks(req *types.SearchTasksRequest) (resp *types.SearchTasksReply, err error) {
	tasks, err := l.svcCtx.MysqlClient.GetTasksByGraph(l.ctx, req.GraphID)
	if err != nil {
		return nil, err
	}
	resp = &types.SearchTasksReply{
		Tasks: make([]*types.Task, 0),
	}
	// reverse
	for i := range tasks {
		task := tasks[len(tasks)-1-i]
		resp.Tasks = append(resp.Tasks, &types.Task{
			Id:         strconv.FormatInt(task.ID, 10),
			Idf:        task.Idf,
			CreateTime: task.CreateTime,
			UpdateTime: task.UpdateTime,
			Status:     task.Status,
			Req:        task.Params,
			Res:        task.Result,
		})
	}
	return resp, nil
}
