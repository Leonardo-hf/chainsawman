package handler

import (
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"

	set "github.com/deckarep/golang-set"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type GetGraph struct {
}

func (h *GetGraph) Handle(params string, taskID int64) (string, error) {
	req := &types.SearchRequest{}
	resp := &types.SearchGraphDetailReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
	}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	nodes, err := config.NebulaClient.GetNodes(req.GraphID, req.Min)
	if err != nil {
		return "", err
	}
	nodeSet := set.NewThreadUnsafeSet()
	for _, node := range nodes {
		resp.Nodes = append(resp.Nodes, &types.Node{
			ID:   node.ID,
			Name: node.Name,
			Desc: node.Desc,
			Deg:  node.Deg,
		})
		nodeSet.Add(node.Name)
	}
	edges, err := config.NebulaClient.GetEdges(req.GraphID)
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
