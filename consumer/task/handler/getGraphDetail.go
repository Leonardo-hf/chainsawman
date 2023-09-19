package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	set "github.com/deckarep/golang-set/v2"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type GetGraphDetail struct {
}

func (h *GetGraphDetail) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.GetGraphDetailRequest{}
	resp := &types.GetGraphDetailReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		NodePacks: make([]*types.NodePack, 0),
		EdgePacks: make([]*types.EdgePack, 0),
	}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	// 获得边
	edges, err := config.NebulaClient.GoFromTopNodes(req.GraphID, req.Top, common.DirectionBoth, req.Max)
	if err != nil {
		return "", err
	}
	// 获得边两端节点id
	nodeSet := set.NewThreadUnsafeSet[int64]()
	for tag, es := range edges {
		resp.EdgePacks = append(resp.EdgePacks, &types.EdgePack{
			Tag:   tag,
			Edges: es,
		})
		for _, e := range es {
			nodeSet.Add(e.Source)
			nodeSet.Add(e.Target)
		}
	}
	// 获得节点
	nodes, err := config.NebulaClient.GetNodesByIds(req.GraphID, nodeSet.ToSlice())
	if err != nil {
		return "", err
	}
	for tag, ns := range nodes {
		resp.NodePacks = append(resp.NodePacks, &types.NodePack{
			Tag:   tag,
			Nodes: ns,
		})
	}
	return jsonx.MarshalToString(resp)
}
