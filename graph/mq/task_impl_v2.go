package mq

import (
	"chainsawman/common"
	"chainsawman/graph/model"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type AsynqImpl struct {
	client    *asynq.Client
	inspector *asynq.Inspector
	config    *AsynqConfig
}

type AsynqConfig struct {
	Addr string
}

type logAdapter struct{}

func (l *logAdapter) Debug(args ...interface{}) {
	logx.Debug(args)
}

func (l *logAdapter) Info(args ...interface{}) {
	logx.Info(args)
}

func (l *logAdapter) Warn(args ...interface{}) {
	logx.Error(args)
}

func (l *logAdapter) Error(args ...interface{}) {
	logx.Error(args)
}

func (l *logAdapter) Fatal(args ...interface{}) {
	logx.Severe(args)
	panic(args)
}

func InitTaskMqV2(cfg *AsynqConfig) TaskMq {
	opt := asynq.RedisClientOpt{Addr: cfg.Addr}
	return &AsynqImpl{
		client:    asynq.NewClient(opt),
		inspector: asynq.NewInspector(opt),
		config:    cfg,
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

func (a *AsynqImpl) ScheduleTask(_ context.Context, task *model.KVTask, cron string) (string, error) {
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: a.config.Addr},
		&asynq.SchedulerOpts{
			Logger: &logAdapter{},
		})
	t := common.TaskIdf(task.Idf)
	content, _ := proto.Marshal(task)
	id, err := scheduler.Register(cron, asynq.NewTask(task.Idf, content), asynq.MaxRetry(3), asynq.Queue(t.Queue))
	if err != nil {
		logx.Error("[Graph] fail to start scheduler, err: %v", err)
		return "", err
	}
	if err = scheduler.Run(); err != nil {
		logx.Error("[Graph] fail to start scheduler, err: %v", err)
		return "", err
	}
	logx.Infof("[Graph] start scheduler, id: %v", id)
	return id, nil
}
