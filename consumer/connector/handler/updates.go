package handler

import (
	msg "chainsawman/client/go"
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/model"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type Updates struct {
}

func (Updates) Handle(m *msg.Msg) error {
	graphID := m.GraphID
	update := &msg.EdgesBody{}
	err := jsonx.UnmarshalFromString(m.Body, update)
	if err != nil {
		return err
	}
	edges := update.Edges
	for _, edge := range edges {
		source, newEdges := edge.Source, edge.Target
		oldEdges, err := config.NebulaClient.GetOutNeighbors(graphID, source)
		if err != nil {
			return err
		}
		for _, target := range oldEdges {
			if id := ifIn(target, newEdges); id == -1 {
				_, err = config.NebulaClient.DeleteEdge(graphID, &model.NebulaEdge{Source: source, Target: target})
				if err != nil {
					return err
				}
			} else {
				_, err = config.NebulaClient.InsertEdge(graphID, &model.NebulaEdge{Source: source, Target: target})
				if err != nil {
					return err
				}
				_, err := config.NebulaClient.AddDeg(graphID, target)
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
