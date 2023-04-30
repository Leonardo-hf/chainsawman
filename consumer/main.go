package main

import (
	"chainsawman/common"
	"chainsawman/consumer/config"
	"chainsawman/consumer/handler"
	"chainsawman/consumer/model"
	"chainsawman/consumer/types/rpc/algo"
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
	//res, err := config.AlgoRPC.QueryAlgo(context.Background(), &algo.Empty{})
	//fmt.Println(res)
	//if err != nil {
	//	logx.Error(err)
	//}
	degree, err := config.AlgoRPC.AvgClustering(context.Background(), &algo.BaseReq{GraphID: 30})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(degree)
	consumerID := uuid.New().String()
	ctx := context.Background()
	for true {
		if err := config.RedisClient.ConsumeTaskMsg(ctx, consumerID, handle); err != nil {
			logx.Errorf("[consumer] consumer fail, err: %v", err)
		}
	}
}

func handle(ctx context.Context, task *model.KVTask) error {
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
