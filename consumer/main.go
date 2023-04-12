package main

import (
	"chainsawman/consumer/config"
	"chainsawman/consumer/handler"
	"chainsawman/consumer/model"

	"context"
	"flag"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var handleTable map[string]handler.Handler

func initHandleTable() {
	handleTable = make(map[string]handler.Handler)
	handleTable["GetGraph"] = &handler.GetGraph{}
	handleTable["GetNode"] = &handler.GetNode{}
	handleTable["Upload"] = &handler.Upload{}
}

func main() {
	flag.Parse()
	var configFile = flag.String("f", "consumer/etc/consumer.yaml", "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	config.Init(&c)
	initHandleTable()
	consumerID := uuid.New().String()
	ctx := context.Background()
	for true {
		if err := config.RedisClient.ConsumeTaskMsg(ctx, consumerID, handle); err != nil {
			logx.Errorf("[consumer] consumer fail, err: %v", err)
		}
	}
}

func handle(ctx context.Context, task *model.KVTask) error {
	h, ok := handleTable[task.Name]
	if !ok {
		return fmt.Errorf("no such method, err: name=%v", task.Name)
	}
	res, err := h.Handle(task.Params, task.Id)
	if err != nil {
		return err
	}
	task.Result = res
	task.Status = model.KVTask_Finished
	if err = config.RedisClient.UpsertTask(ctx, task); err != nil {
		return err
	}
	_, err = config.MysqlClient.UpdateTask(&model.Task{
		ID:     task.Id,
		Status: int64(task.Status),
		Result: res,
	})
	return err
}
