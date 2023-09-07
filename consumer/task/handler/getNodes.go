package handler

import (
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"github.com/zeromicro/go-zero/core/jsonx"
	"math"
)

type GetNodes struct {
}

func (h *GetNodes) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.GetNodesRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	resp := &types.GetNodesReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		NodePacks: make([]*types.NodePack, 0),
	}
	nodes, err := config.NebulaClient.GetTopNodes(req.GraphID, math.MaxInt)
	if err != nil {
		return "", err
	}
	for tag, v := range nodes {
		resp.NodePacks = append(resp.NodePacks, &types.NodePack{
			Tag:   tag,
			Nodes: v,
		})
	}
	return jsonx.MarshalToString(resp)
}
