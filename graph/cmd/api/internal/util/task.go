package util

import (
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/model"
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func PublishTask(ctx context.Context, svcCtx *svc.ServiceContext, graphID int64, taskIDf string, req interface{}) (string, error) {
	params, err := jsonx.MarshalToString(req)
	if err != nil {
		return "", err
	}
	taskID := uuid.New().String()
	task := &model.KVTask{
		Id:      taskID,
		Params:  params,
		Status:  model.KVTask_New,
		Idf:     taskIDf,
		GraphID: graphID,
	}
	// 缓存任务
	err = svcCtx.RedisClient.UpsertTask(ctx, task)
	if err != nil {
		return "", err
	}
	// 发布任务
	_, err = svcCtx.TaskMq.ProduceTaskMsg(ctx, task)
	if err != nil {
		return "", err
	}
	return taskID, nil
}

func FetchTask(ctx context.Context, svcCtx *svc.ServiceContext, taskID string, resp interface{}) error {
	result := ""
	status := model.KVTask_New
	task, err := svcCtx.RedisClient.GetTaskById(ctx, taskID)
	if err == nil {
		// redis里有记录
		result = task.Result
		status = task.Status
		if status == model.KVTask_Finished {
			err = jsonx.UnmarshalFromString(result, resp)
		}
	}
	return err
}
