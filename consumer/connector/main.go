package main

import (
	msg "chainsawman/client/go"
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/handler"

	"context"
	"flag"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var handleTable map[msg.Opt]handler.Handler

func initHandleTable() {
	handleTable = make(map[msg.Opt]handler.Handler)
	handleTable[msg.Creates] = &handler.Creates{}
	handleTable[msg.Updates] = &handler.Updates{}
	handleTable[msg.Deletes] = &handler.Deletes{}
}

// TODO: 很多问题没有解决
// TODO: 顺序消费
// TODO: 事务
func main() {
	flag.Parse()
	var configFile = flag.String("f", "consumer/connector/etc/consumer.yaml", "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	config.Init(&c)
	initHandleTable()
	consumerID := uuid.New().String()
	ctx := context.Background()
	for true {
		if err := config.ImportMq.ConsumeImportMsg(ctx, consumerID, handle); err != nil {
			logx.Errorf("[connector] consumer fail, err: %v", err)
		}
	}
}

func handle(_ context.Context, m *msg.Msg) error {
	return handleTable[m.Opt].Handle(m)
}
