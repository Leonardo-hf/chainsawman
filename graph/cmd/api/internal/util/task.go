package util

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/model"
	"github.com/google/uuid"
	"strconv"

	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func PublishTask(ctx context.Context, svcCtx *svc.ServiceContext, graphID int64, taskIDf string, req interface{}) (string, error) {
	params, err := jsonx.MarshalToString(req)
	if err != nil {
		return "", err
	}
	fmt.Println(taskIDf)
	persistTask := &model.Task{
		Params:  params,
		Idf:     taskIDf,
		GraphID: graphID,
	}
	taskID := uuid.New().String()
	tid := ""
	// 持久化任务
	if common.TaskIdf(taskIDf).Persistent {
		err = svcCtx.MysqlClient.InsertTask(ctx, persistTask)
		if err != nil {
			return "", err
		}
		taskID = strconv.FormatInt(persistTask.ID, 10)
		defer func() {
			_, err := svcCtx.MysqlClient.UpdateTaskTIDByID(ctx, persistTask.ID, tid)
			logx.Errorf("[Graph] update task tid fail, err=%v", err)
		}()
	}
	task := &model.KVTask{
		Id:         taskID,
		Params:     params,
		Status:     model.KVTask_New,
		CreateTime: persistTask.CreateTime,
		Idf:        taskIDf,
	}
	// 临时保存任务
	err = svcCtx.RedisClient.UpsertTask(ctx, task)
	if err != nil {
		return "", err
	}
	// 发布任务
	tid, err = svcCtx.TaskMq.ProduceTaskMsg(ctx, task)
	if err != nil {
		return "", err
	}
	return taskID, nil
}

func FetchTask(ctx context.Context, svcCtx *svc.ServiceContext, taskID string, idf string, resp interface{}) error {
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
		taskIDInt, _ := strconv.ParseInt(taskID, 10, 64)
		oTask, err := svcCtx.MysqlClient.GetTaskByID(ctx, taskIDInt)
		if err != nil {
			return err
		}
		status = model.KVTask_Status(oTask.Status)
		result = oTask.Result
		// 更新redis
		err = svcCtx.RedisClient.UpsertTask(ctx, &model.KVTask{
			Id:         strconv.FormatInt(oTask.ID, 10),
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
