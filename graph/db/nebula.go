package db

import "chainsawman/graph/model"

type NebulaClient interface {
	DropGraph(graph int64) error

	CreateGraph(graph int64) error

	GetAllNodes(graph int64) ([]*model.Node, error)
}
