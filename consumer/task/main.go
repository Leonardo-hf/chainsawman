package main

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/handler"
	"chainsawman/consumer/task/model"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"

	"context"
	"flag"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var handleTable map[string]handler.Handler

func initHandleTable() {
	handleTable = make(map[string]handler.Handler)
	handleTable[common.GraphGet] = &handler.GetGraphDetail{}
	handleTable[common.GraphUpdate] = &handler.UpdateGraph{}
	handleTable[common.GraphNeighbors] = &handler.GetNeighbors{}
	handleTable[common.GraphNodes] = &handler.GetNodes{}
	handleTable[common.GraphCreate] = &handler.CreateGraph{}
	handleTable[common.AlgoExec] = &handler.AlgoExec{}
}

func main() {
	flag.Parse()
	defaultCfg := "consumer/task/etc/consumer.yaml"
	switch os.Getenv("CHS_ENV") {
	case "docker-compose":
		defaultCfg = "consumer/task/etc/consumer-docker.yaml"
	case "pre":
		defaultCfg = "consumer/task/etc/consumer-pre.yaml"
	}
	var configFile = flag.String("f", defaultCfg, "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	config.Init(&c)
	initHandleTable()
	if c.IsTaskV2Enabled() {
		reportError := func(ctx context.Context, task *asynq.Task, err error) {
			retried, _ := asynq.GetRetryCount(ctx)
			maxRetry, _ := asynq.GetMaxRetry(ctx)
			if retried >= maxRetry {
				logx.Errorf("retry exhausted for task %s: %w", task.Type, err)
			}
		}
		srv := asynq.NewServer(
			asynq.RedisClientOpt{Addr: c.TaskMq.Addr},
			asynq.Config{
				Concurrency: 2,
				Queues: map[string]int{
					common.PHigh:   6,
					common.PMedium: 3,
					common.PLow:    1,
				},
				ErrorHandler: asynq.ErrorHandlerFunc(reportError),
			},
		)
		mux := asynq.NewServeMux()
		for idf, h := range handleTable {
			mux.HandleFunc(idf, getAsynqHandler(h))
		}
		if err := srv.Run(mux); err != nil {
			logx.Errorf("could not run server: %v", err)
			panic(err)
		}
		return
	}
	consumerID := uuid.New().String()
	ctx := context.Background()
	for true {
		if err := config.TaskMq.ConsumeTaskMsg(ctx, consumerID, getRedisHandler()); err != nil {
			logx.Errorf("[task] consumer fail, err: %v", err)
		}
	}
}

func handle(ctx context.Context, task *model.KVTask, h handler.Handler) error {
	res, err := h.Handle(task)
	if err == config.DelayTaskErr {
		logx.Infof("[Task] %v", err)
		return nil
	} else if err != nil {
		fmt.Println(err)
		return err
	}
	logx.Infof("[Task] finish task, idf=%v", task.Idf)
	task.Result = res
	task.Status = model.KVTask_Finished
	task.UpdateTime = time.Now().UTC().Unix()
	if err = config.RedisClient.UpsertTask(ctx, task); err != nil {
		return err
	}
	if common.TaskIdf(task.Idf).Persistent {
		taskIDInt, _ := strconv.ParseInt(task.Id, 10, 64)
		_, err = config.MysqlClient.UpdateTaskByID(ctx, &model.Task{
			ID:     taskIDInt,
			Status: int64(task.Status),
			Result: res,
		})
	}
	return err
}

func getRedisHandler() func(ctx context.Context, task *model.KVTask) error {
	return func(ctx context.Context, t *model.KVTask) error {
		h, ok := handleTable[t.Idf]
		if !ok {
			return fmt.Errorf("no such method, err: idf=%v", common.TaskIdf(t.Idf))
		}
		return handle(ctx, t, h)
	}
}

func getAsynqHandler(h handler.Handler) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		t := &model.KVTask{}
		_ = proto.Unmarshal(task.Payload(), t)
		return handle(ctx, t, h)
	}
}
