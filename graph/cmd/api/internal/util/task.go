package util

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/model"
	"fmt"

	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func PublishTask(ctx context.Context, svcCtx *svc.ServiceContext, graphID int64, taskIDf string, req interface{}) (int64, error) {
	params, err := jsonx.MarshalToString(req)
	if err != nil {
		return 0, err
	}
	fmt.Println(taskIDf)
	task := &model.Task{
		Params:  params,
		Idf:     taskIDf,
		GraphID: graphID,
	}
	// 保存任务
	err = svcCtx.MysqlClient.InsertTask(ctx, task)
	if err != nil {
		return 0, err
	}
	// 发布任务

	tid, err := svcCtx.TaskMq.ProduceTaskMsg(ctx, &model.KVTask{
		Id:         task.ID,
		Params:     task.Params,
		Status:     model.KVTask_New,
		CreateTime: task.CreateTime,
		UpdateTime: task.UpdateTime,
		Idf:        taskIDf,
	})
	if err != nil {
		return 0, err
	}
	// TODO: 更新任务TID，这样似乎不好
	_, err = svcCtx.MysqlClient.UpdateTaskTIDByID(ctx, task.ID, tid)
	if err != nil {
		return 0, err
	}
	return task.ID, nil
}

func FetchTask(ctx context.Context, svcCtx *svc.ServiceContext, taskID int64, idf string, resp interface{}) error {
	result := ""
	status := model.KVTask_New
	task, err := svcCtx.RedisClient.GetTaskById(ctx, taskID)
	if err == nil {
		// redis里有记录
		result = task.Result
		status = task.Status
	} else if err == redis.Nil {
		// 如果没有持久化则直接返回
		if !common.TaskIdf(idf).Persistent {
			return nil
		}
		// redis里没有记录，查询mysql
		oTask, err := svcCtx.MysqlClient.GetTaskByID(ctx, taskID)
		if err != nil {
			return err
		}
		status = model.KVTask_Status(oTask.Status)
		result = oTask.Result
		// 更新redis
		err = svcCtx.RedisClient.UpsertTask(ctx, &model.KVTask{
			Id:         oTask.ID,
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
