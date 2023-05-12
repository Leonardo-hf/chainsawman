package handler

import (
	msg "chainsawman/client/go"
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/model"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type Deletes struct {
}

// Handle 忽略对度数的影响
func (Deletes) Handle(m *msg.Msg) error {
	graphId := m.GraphID
	de := &msg.NodesBody{}
	err := jsonx.UnmarshalFromString(m.Body, de)
	if err != nil {
		return err
	}
	for _, n := range de.Nodes {
		_, err := config.NebulaClient.DeleteNode(graphId, n.NodeID)
		if err != nil {
			return err
		}
		neighbors, err := config.NebulaClient.GetOutNeighbors(graphId, n.NodeID)
		if err != nil {
			return err
		}
		for _, target := range neighbors {
			_, err := config.NebulaClient.DeleteEdge(graphId, &model.NebulaEdge{
				Source: n.NodeID,
				Target: target,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
