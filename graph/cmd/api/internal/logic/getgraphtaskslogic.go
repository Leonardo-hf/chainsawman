package logic

import (
	"chainsawman/common"
	"context"

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
	resp = &types.SearchTasksReply{}
	// reverse
	for i := range tasks {
		task := tasks[len(tasks)-1-i]
		idf := common.TaskIdf(task.Idf)
		resp.Tasks = append(resp.Tasks, &types.Task{
			Id:         task.ID,
			Idf:        task.Idf,
			Desc:       idf.Desc(),
			CreateTime: task.CreateTime,
			UpdateTime: task.UpdateTime,
			Status:     task.Status,
			Req:        idf.Params(task.Params),
			Res:        idf.Res(task.Result),
		})
	}
	return resp, nil
}
