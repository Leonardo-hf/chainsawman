package main

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/handler"
	"chainsawman/consumer/task/model"
	"github.com/google/uuid"

	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var handleTable map[common.TaskIdf]handler.Handler

func initHandleTable() {
	handleTable = make(map[common.TaskIdf]handler.Handler)
	handleTable[common.GraphGet] = &handler.GetGraph{}
	handleTable[common.GraphNeighbors] = &handler.GetNeighbors{}
	handleTable[common.GraphCreate] = &handler.Upload{}
	handleTable[common.AlgoDegree] = &handler.AlgoDegree{}
	handleTable[common.AlgoPagerank] = &handler.AlgoPageRank{}
	handleTable[common.AlgoVoterank] = &handler.AlgoVoteRank{}
	handleTable[common.AlgoCloseness] = &handler.AlgoCloseness{}
	handleTable[common.AlgoBetweenness] = &handler.AlgoBetweenness{}
	handleTable[common.AlgoAvgCC] = &handler.AlgoAvgCC{}
	handleTable[common.AlgoLouvain] = &handler.AlgoLouvain{}
}

func main() {
	flag.Parse()
	var configFile = flag.String("f", "consumer/task/etc/consumer.yaml", "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	config.Init(&c)
	initHandleTable()
	consumerID := uuid.New().String()
	ctx := context.Background()
	for true {
		if err := config.TaskMq.ConsumeTaskMsg(ctx, consumerID, handle); err != nil {
			logx.Errorf("[task] consumer fail, err: %v", err)
		}
	}
}

func handle(ctx context.Context, task *model.KVTask) error {
	fmt.Println(common.TaskIdf(task.Idf))
	h, ok := handleTable[common.TaskIdf(task.Idf)]
	if !ok {
		return fmt.Errorf("no such method, err: idf=%v", common.TaskIdf(task.Idf).Desc())
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
	_, err = config.MysqlClient.UpdateTaskByID(&model.Task{
		ID:     task.Id,
		Status: int64(task.Status),
		Result: res,
	})
	return err
}
