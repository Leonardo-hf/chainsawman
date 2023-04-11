package handler

import (
	"chainsawman/consumer/config"
	"chainsawman/consumer/types"

	set "github.com/deckarep/golang-set"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type GetGraph struct {
}

func (h *GetGraph) Handle(params string) (string, error) {
	req := &types.SearchRequest{}
	resp := &types.SearchGraphDetailReply{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	nodes, err := config.NebulaClient.GetNodes(req.Graph, req.Min)
	if err != nil {
		return "", err
	}
	nodeSet := set.NewThreadUnsafeSet()
	for _, node := range nodes {
		resp.Nodes = append(resp.Nodes, &types.Node{
			Name: node.Name,
			Desc: node.Desc,
		})
		nodeSet.Add(node.Name)
	}
	edges, err := config.NebulaClient.GetEdges(req.Graph)
	if err != nil {
		return "", err
	}
	for _, edge := range edges {
		if nodeSet.Contains(edge.Source) || nodeSet.Contains(edge.Target) {
			resp.Edges = append(resp.Edges, &types.Edge{
				Source: edge.Source,
				Target: edge.Target,
			})
		}
	}
	return jsonx.MarshalToString(resp)
}
