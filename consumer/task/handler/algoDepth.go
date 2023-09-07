package handler

import (
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"chainsawman/consumer/task/types/rpc/algo"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoDepth struct {
}

func (h *AlgoDepth) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.AlgoVoteRankRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	edgeTags, err := PrepareForAlgoRPC(ctx, req.GraphID)
	if err != nil {
		return "", err
	}
	res, err := config.AlgoRPC.Depth(ctx, &algo.BaseReq{GraphID: req.GraphID, EdgeTags: edgeTags})
	if err != nil {
		return "", err
	}
	return HandleAlgoRPCRes(res, req.GraphID, taskID)
}
