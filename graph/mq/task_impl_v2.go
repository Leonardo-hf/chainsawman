package mq

import (
	"chainsawman/common"
	"chainsawman/graph/model"
	"github.com/golang/protobuf/proto"

	"context"

	"github.com/hibiken/asynq"
)

type AsynqImpl struct {
	client    *asynq.Client
	inspector *asynq.Inspector
}

type AsynqConfig struct {
	Addr string
}

func InitTaskMqV2(cfg *AsynqConfig) TaskMq {
	opt := asynq.RedisClientOpt{Addr: cfg.Addr}
	return &AsynqImpl{
		client:    asynq.NewClient(opt),
		inspector: asynq.NewInspector(opt),
	}
}

func (a *AsynqImpl) enqueue(ctx context.Context, idf string, task *model.KVTask) (*asynq.TaskInfo, error) {
	t := common.TaskIdf(idf)
	content, _ := proto.Marshal(task)
	// TODO: maxRetry
	return a.client.EnqueueContext(ctx, asynq.NewTask(idf, content), asynq.MaxRetry(0), asynq.Queue(t.Queue))
}

func (a *AsynqImpl) ProduceTaskMsg(ctx context.Context, task *model.KVTask) (string, error) {
	info, err := a.enqueue(ctx, task.Idf, task)
	if err != nil {
		return "", err
	}
	return info.ID, nil
}

func (a *AsynqImpl) DelTaskMsg(_ context.Context, task *model.KVTask) error {
	t := common.TaskIdf(task.Idf)
	return a.inspector.DeleteTask(t.Queue, task.Tid)
}
