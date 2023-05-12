package handler

import (
	msg "chainsawman/client/go"
	"chainsawman/consumer/connector/config"
	"chainsawman/consumer/connector/model"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type Creates struct {
}

func (Creates) Handle(m *msg.Msg) error {
	graphId := m.GraphID
	create := &msg.NodesBody{}
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
