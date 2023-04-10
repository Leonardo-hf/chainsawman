package db

import (
	"chainsawman/graph/model"
)

// NebulaClient TODO: 删掉用不上的接口
type NebulaClient interface {
	CreateGraph(graph string) error

	InsertNode(graph string, node *model.Node) (int, error)

	MultiInsertNodes(graph string, nodes []*model.Node) (int, error)

	InsertEdge(graph string, edge *model.Edge) (int, error)

	MultiInsertEdges(graph string, edges []*model.Edge) (int, error)

	GetGraphs() ([]*model.Graph, error)

	GetNodes(graph string, min int64) ([]*model.Node, error)

	GetEdges(graph string) ([]*model.Edge, error)

	DropGraph(graph string) error
}
