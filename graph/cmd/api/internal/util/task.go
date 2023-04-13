package util

import (
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/model"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func PublishTask(ctx context.Context, svcCtx *svc.ServiceContext, taskName string, req interface{}) (int64, error) {
	params, err := jsonx.MarshalToString(req)
	if err != nil {
		return 0, err
	}
	task := &model.Task{
		Params: params,
		Name:   taskName,
	}
	// 保存任务
	err = svcCtx.MysqlClient.InsertTask(task)
	if err != nil {
		return 0, err
	}
	// 发布任务
	err = svcCtx.RedisClient.ProduceTaskMsg(ctx, &model.KVTask{
		Id:         task.ID,
		Name:       task.Name,
		Params:     task.Params,
		Status:     model.KVTask_Status(task.Status),
		CreateTime: task.CreateTime,
		UpdateTime: task.UpdateTime,
	})
	if err != nil {
		return 0, err
	}
	return task.ID, nil
}

func FetchTask(ctx context.Context, svcCtx *svc.ServiceContext, taskID int64, resp interface{}) error {
	result := ""
	status := model.KVTask_New
	task, err := svcCtx.RedisClient.GetTaskById(ctx, taskID)
	if err == nil {
		// redis里有记录
		result = task.Result
		status = task.Status
	} else if err == redis.Nil {
		// redis里没有记录，查询mysql
		oTask, err := svcCtx.MysqlClient.SearchTaskById(taskID)
		if err != nil {
			return err
		}
		status = model.KVTask_Status(oTask.Status)
		result = oTask.Result
		// 更新redis
		err = svcCtx.RedisClient.UpsertTask(ctx, &model.KVTask{
			Id:         oTask.ID,
			Name:       oTask.Name,
			Params:     oTask.Params,
			Status:     model.KVTask_Status(oTask.Status),
			CreateTime: oTask.CreateTime,
			UpdateTime: oTask.UpdateTime,
			Result:     oTask.Result,
		})
		if err != nil {
			return err
		}
	} else {
		return err
	}
	if status == model.KVTask_Finished {
		if err = jsonx.UnmarshalFromString(result, resp); err != nil {
			return err
		}
	}
	return nil
}
