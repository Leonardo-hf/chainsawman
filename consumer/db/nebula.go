package db

import (
	"chainsawman/consumer/model"
)

type NebulaClient interface {
	CreateGraph(graph int64) error

	InsertNode(graph int64, node *model.Node) (int, error)

	MultiInsertNodes(graph int64, nodes []*model.Node) (int, error)

	InsertEdge(graph int64, edge *model.Edge) (int, error)

	MultiInsertEdges(graph int64, edges []*model.Edge) (int, error)

	GetNodes(graph int64, min int64) ([]*model.Node, error)

	GetEdges(graph int64) ([]*model.Edge, error)

	GetNeighbors(graph int64, nodeID int64, min int64, distance int64) ([]*model.Node, []*model.Edge, error)

	DropGraph(graph int64) error
}
