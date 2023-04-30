package handler

import (
	"chainsawman/consumer/config"
	"chainsawman/consumer/model"
	"chainsawman/consumer/types"

	"github.com/zeromicro/go-zero/core/jsonx"
)

type GetNeighbors struct {
}

func (h *GetNeighbors) Handle(params string, taskID int64) (string, error) {
	req := &types.SearchNodeRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	nodes, edges, err := config.NebulaClient.GetNeighbors(req.GraphID, req.NodeID, req.Min, req.Distance)
	if err != nil {
		return "", err
	}
	// TODO: 也许可以单独写一些类型转化的Format方法
	var nodesRet []*types.Node
	var edgesRet []*types.Edge
	for _, node := range nodes {
		nodesRet = append(nodesRet, &types.Node{
			ID:   node.ID,
			Name: node.Name,
			Desc: node.Desc,
			Deg:  node.Deg,
		})
	}
	for _, edge := range edges {
		edgesRet = append(edgesRet, &types.Edge{
			Source: edge.Source,
			Target: edge.Target,
		})
	}
	resp := &types.SearchNodeReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Info:  nodesRet[0],
		Nodes: nodesRet,
		Edges: edgesRet,
	}
	return jsonx.MarshalToString(resp)
}
