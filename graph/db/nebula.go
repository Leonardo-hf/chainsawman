package db

import "chainsawman/graph/model"

type NebulaClient interface {
	MatchNodes(graph int64, keywords string, group *model.Group) (map[string][]*MatchNode, error)
	MatchNodesByTag(graph int64, keywords string, node *model.Node) ([]*MatchNode, error)
	DropGraph(graph int64) error
}

type MatchNode struct {
	Id          int64  `json:"id"`
	PrimaryAttr string `json:"primaryAttr"`
}
