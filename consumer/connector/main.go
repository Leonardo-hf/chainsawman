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

// TODO: 很多问题没有解决
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
	update := &msg.UpdateBody{}
	err := jsonx.UnmarshalFromString(m.Body, update)
	if err != nil {
		return err
	}
	return handleUpdate(ctx, update)
	//switch m.Entity {
	//case msg.Node:
	//	node := &msg.NodeBody{}
	//	err := jsonx.UnmarshalFromString(m.Body, node)
	//	if err != nil {
	//		return err
	//	}
	//	return handleNode(ctx, node, m.GraphID, m.Opt)
	//case msg.Edge:
	//	node := &msg.EdgeBody{}
	//	err := jsonx.UnmarshalFromString(m.Body, node)
	//	if err != nil {
	//		return err
	//	}
	//	return handleEdge(ctx, node, m.GraphID, m.Opt)
	//}
	//return fmt.Errorf("[connector] wrong entity flag, value = %v", m.Entity)
}

func handleUpdate(_ context.Context, update *msg.UpdateBody) error {
	graphId := update.GraphId
	edges := update.Edges
	for i, newEdge := range edges {
		oldEdge, err := config.NebulaClient.GetOutNeighbors(graphId, i)
		if err != nil {
			return err
		}
		flags := make([]bool, len(newEdge))
		for _, old := range oldEdge {
			if id := ifIn(old, newEdge); id == -1 {
				_, err = config.NebulaClient.DeleteEdge(graphId, &msg.EdgeBody{Source: i, Target: old})
				if err != nil {
					return err
				}
			} else {
				flags[id] = true
			}
		}
		for index, f := range flags {
			if !f {
				_, err = config.NebulaClient.InsertEdge(graphId, &msg.EdgeBody{Source: i, Target: newEdge[index]})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func ifIn(old int64, newEDge []int64) int64 {
	for i := range newEDge {
		if old == newEDge[i] {
			return int64(i)
		}
	}
	return -1
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
