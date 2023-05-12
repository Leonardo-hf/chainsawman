package handler

import (
	msg "chainsawman/client/go"
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/model"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type DeleteEdges struct {
}

// Handle 忽略对度数的影响
func (DeleteEdges) Handle(m *msg.Msg) error {
	graphID := m.GraphID
	body := &msg.SplitEdgesBody{}
	err := jsonx.UnmarshalFromString(m.Body, body)
	if err != nil {
		return err
	}
	for _, edge := range body.Edges {
		_, err = config.NebulaClient.DeleteEdge(graphID, &model.NebulaEdge{
			Source: edge[0],
			Target: edge[1],
		})
		if err != nil {
			return err
		}
	}
	return nil
}
