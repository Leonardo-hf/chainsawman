package db

import (
	"chainsawman/graph/model"
)

type NebulaClient interface {
	CreateGraph(graph string) error

	InsertNode(graph string, node *model.Node) (int, error)

	MultiInsertNodes(graph string, nodes []*model.Node) (int, error)

	InsertEdge(graph string, edge *model.Edge) (int, error)

	MultiInsertEdges(graph string, edges []*model.Edge) (int, error)

	GetGraphs() ([]*model.Graph, error)

	GetNodes(graph string) ([]*model.Node, error)

	GetEdges(graph string) ([]*model.Edge, error)

	DropGraph(graph string) error
}
