package db

import (
	"chainsawman/consumer/connector/msg"
)

type NebulaClient interface {
	InsertNode(graphID int64, node *msg.NodeBody) (int, error)

	InsertEdge(graphID int64, edge *msg.EdgeBody) (int, error)

	UpdateNode(graphID int64, node *msg.NodeBody) (int, error)

	DeleteNode(graphID int64, nodeID int64) (int, error)

	DeleteEdge(graphID int64, edge *msg.EdgeBody) (int, error)

	GetOutNeighbors(graph int64, nodeID int64) ([]int64, error)
}
