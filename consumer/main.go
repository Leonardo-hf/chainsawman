package main

import (
	"chainsawman/consumer/config"
	"chainsawman/consumer/handler"
	"chainsawman/consumer/model"
	"chainsawman/consumer/types/rpc/algo"

	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var handleTable map[string]handler.Handler

func initHandleTable() {
	handleTable = make(map[string]handler.Handler)
	handleTable["GetGraph"] = &handler.GetGraph{}
	handleTable["GetNeighbors"] = &handler.GetNeighbors{}
	handleTable["Upload"] = &handler.Upload{}
	handleTable["AlgoDegree"] = &handler.AlgoDegree{}
	handleTable["AlgoPageRank"] = &handler.AlgoPageRank{}
}

func main() {
	flag.Parse()
	var configFile = flag.String("f", "consumer/etc/consumer.yaml", "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	config.Init(&c)
	initHandleTable()
	//res, err := config.AlgoRPC.CreateAlgo(context.Background(), &algo.CreateAlgoReq{
	//	Name: "testname",
	//	Desc: "testdesc",
	//	Params: []*algo.Element{
	//		{
	//			Key:  "123",
	//			Type: algo.Element_INT64,
	//		},
	//		{
	//			Key:  "456",
	//			Type: algo.Element_LIST_DOUBLE,
	//		},
	//	},
	//	Type: algo.Algo_Rank,
	//})
	res, err := config.AlgoRPC.QueryAlgo(context.Background(), &algo.Empty{})
	fmt.Println(res)
	if err != nil {
		logx.Error(err)
	}
	//consumerID := uuid.New().String()
	//ctx := context.Background()
	//for true {
	//	if err := config.RedisClient.ConsumeTaskMsg(ctx, consumerID, handle); err != nil {
	//		logx.Errorf("[consumer] consumer fail, err: %v", err)
	//	}
	//}
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
	_, err = config.MysqlClient.UpdateTaskByID(&model.Task{
		ID:     task.Id,
		Status: int64(task.Status),
		Result: res,
	})
	return err
}
