package db

import (
	"chainsawman/graph/model"
)

// NebulaClient TODO: 删掉用不上的接口
type NebulaClient interface {
	GetGraphs() ([]*model.Graph, error)

	DropGraph(graph string) error
}
