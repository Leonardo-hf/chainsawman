package main

import (
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/model"
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

func handle(ctx context.Context, m *model.Msg) error {
	switch m.Opt {
	case model.Creates:
		return handleCreate(ctx, m)
	case model.Updates:
		return handleUpdate(ctx, m)
	case model.Deletes:
		return handleDelete(ctx, m)
	}
	return fmt.Errorf("[connector] wrong OPT flag, value = %v", m.Opt)
}

// TODO: 度数问题
func handleUpdate(_ context.Context, m *model.Msg) error {
	graphId := m.GraphID
	update := &model.EdgesBody{}
	err := jsonx.UnmarshalFromString(m.Body, update)
	if err != nil {
		return err
	}
	edges := update.Edges
	for _, edge := range edges {
		source, newEdges := edge.Source, edge.Target
		oldEdges, err := config.NebulaClient.GetOutNeighbors(graphId, source)
		if err != nil {
			return err
		}
		for _, target := range oldEdges {
			if id := ifIn(target, newEdges); id == -1 {
				_, err = config.NebulaClient.DeleteEdge(graphId, &model.NebulaEdge{Source: source, Target: target})
				if err != nil {
					return err
				}
			} else {
				_, err = config.NebulaClient.InsertEdge(graphId, &model.NebulaEdge{Source: source, Target: target})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func handleCreate(_ context.Context, m *model.Msg) error {
	graphId := m.GraphID
	create := &model.NodesBody{}
	err := jsonx.UnmarshalFromString(m.Body, create)
	if err != nil {
		return err
	}
	nodes := create.Nodes
	nebulaNodes := make([]*model.NebulaNode, len(nodes))
	for _, n := range nodes {
		nebulaNodes = append(nebulaNodes, &model.NebulaNode{
			ID:   n.NodeID,
			Name: n.Name,
			Desc: n.Desc,
			Deg:  0,
		})
	}
	_, err = config.NebulaClient.MultiInsertNodes(graphId, nebulaNodes)
	if err != nil {
		return err
	}
	return nil
}

// TODO: 删除边与度数问题
func handleDelete(_ context.Context, m *model.Msg) error {
	graphId := m.GraphID
	de := &model.NodesBody{}
	err := jsonx.UnmarshalFromString(m.Body, de)
	if err != nil {
		return err
	}
	for _, n := range de.Nodes {
		_, err := config.NebulaClient.DeleteNode(graphId, n.NodeID)
		if err != nil {
			return err
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
