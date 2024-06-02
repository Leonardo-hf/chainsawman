package util

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/model"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
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

func CreateCronTask(ctx context.Context, svcCtx *svc.ServiceContext, taskIDf string) (string, error) {
	taskID := uuid.New().String()
	task := &model.KVTask{
		Id:     taskID,
		Status: model.KVTask_New,
		Idf:    taskIDf,
	}
	// 判断幂等
	week := 7 * 24 * time.Hour
	key := fmt.Sprintf("%v_%v", taskIDf, time.Now().Format("2006-01-02/15:04"))
	idempotent, err := svcCtx.RedisClient.CheckIdempotent(ctx, key, week)
	if err != nil {
		return "", err
	}
	if idempotent {
		t := common.TaskIdf(taskIDf)
		// 发布任务
		_, err = svcCtx.TaskMq.ScheduleTask(ctx, task, t.Cron)
		if err != nil {
			return "", err
		}
		return taskID, nil
	}
	return "", nil
}
