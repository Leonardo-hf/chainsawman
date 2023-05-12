package handler

import (
	msg "chainsawman/client/go"
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/model"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type InsertEdges struct {
}

func (InsertEdges) Handle(m *msg.Msg) error {
	graphID := m.GraphID
	body := &msg.SplitEdgesBody{}
	err := jsonx.UnmarshalFromString(m.Body, body)
	if err != nil {
		return err
	}
	edges := make([]*model.NebulaEdge, 0)
	for _, edge := range body.Edges {
		edges = append(edges, &model.NebulaEdge{
			Source: edge[0],
			Target: edge[1],
		})
		for _, nodeID := range edge {
			_, err := config.NebulaClient.AddDeg(graphID, nodeID)
			if err != nil {
				return err
			}
		}
	}
	_, err = config.NebulaClient.MultiInsertEdges(graphID, edges)
	if err != nil {
		return err
	}
	return nil
}
