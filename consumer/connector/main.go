package main

import (
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/msg"

	"context"
	"flag"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	flag.Parse()
	var configFile = flag.String("f", "consumer/connector/etc/consumer.yaml", "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	config.Init(&c)
	consumerID := uuid.New().String()
	ctx := context.Background()
	for true {
		if err := config.ImportMq.ConsumeImportMsg(ctx, consumerID, handle); err != nil {
			logx.Errorf("[connector] consumer fail, err: %v", err)
		}
	}
}

func handle(ctx context.Context, m *msg.Msg) error {
	switch m.Entity {
	case msg.Node:
		node := &msg.NodeBody{}
		err := jsonx.UnmarshalFromString(m.Body, node)
		if err != nil {
			return err
		}
		return handleNode(ctx, node, m.GraphID, m.Opt)
	case msg.Edge:
		node := &msg.EdgeBody{}
		err := jsonx.UnmarshalFromString(m.Body, node)
		if err != nil {
			return err
		}
		return handleEdge(ctx, node, m.GraphID, m.Opt)
	}
	return fmt.Errorf("[connector] wrong entity flag, value = %v", m.Entity)
}

func handleEdge(_ context.Context, edge *msg.EdgeBody, graphID int64, opt msg.OptFlag) error {
	switch opt {
	case msg.Create:
		_, err := config.NebulaClient.InsertEdge(graphID, edge)
		if err != nil {
			return err
		}
	case msg.Delete:
		_, err := config.NebulaClient.DeleteEdge(graphID, edge)
		if err != nil {
			return err
		}
	}
	return fmt.Errorf("[connector] wrong opt flag, value = %v", opt)
}

func handleNode(_ context.Context, node *msg.NodeBody, graphID int64, opt msg.OptFlag) error {
	switch opt {
	case msg.Create:
		_, err := config.NebulaClient.InsertNode(graphID, node)
		if err != nil {
			return err
		}
	case msg.Update:
		_, err := config.NebulaClient.UpdateNode(graphID, node)
		if err != nil {
			return err
		}
	case msg.Delete:
		_, err := config.NebulaClient.DeleteNode(graphID, node.ID)
		if err != nil {
			return err
		}
	}
	return fmt.Errorf("[connector] wrong opt flag, value = %v", opt)
}
