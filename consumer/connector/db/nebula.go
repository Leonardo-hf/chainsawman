package db

import (
	"chainsawman/consumer/connector/model"
)

type NebulaClient interface {
	MultiInsertNodes(graph int64, nodes []*model.NebulaNode) (int, error)

	MultiInsertEdges(graph int64, edges []*model.NebulaEdge) (int, error)

	InsertEdge(graphID int64, edge *model.NebulaEdge) (int, error)

	AddDeg(graphID int64, nodeID int64) (int, error)

	UpdateNode(graphID int64, node *model.NebulaNode) (int, error)

	DeleteNode(graphID int64, nodeID int64) (int, error)

	DeleteEdge(graphID int64, edge *model.NebulaEdge) (int, error)

	GetOutNeighbors(graph int64, nodeID int64) ([]int64, error)
}
