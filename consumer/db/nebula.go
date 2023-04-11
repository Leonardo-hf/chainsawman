package db

import (
	"chainsawman/consumer/model"
)

type NebulaClient interface {
	CreateGraph(graph string) error

	InsertNode(graph string, node *model.Node) (int, error)

	MultiInsertNodes(graph string, nodes []*model.Node) (int, error)

	InsertEdge(graph string, edge *model.Edge) (int, error)

	MultiInsertEdges(graph string, edges []*model.Edge) (int, error)

	GetNodeByName(graph string, name string) (*model.Node, error)

	GetNodes(graph string, min int64) ([]*model.Node, error)

	GetEdges(graph string) ([]*model.Edge, error)

	GetNeighbors(graph string, node string, min int64, distance int64) ([]*model.Node, []*model.Edge, error)

	DropGraph(graph string) error
}
