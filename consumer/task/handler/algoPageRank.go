package handler

import (
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"chainsawman/consumer/task/types/rpc/algo"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoPageRank struct {
}

func (h *AlgoPageRank) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.AlgoPageRankRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	edgeTags, err := PrepareForAlgoRPC(ctx, req.GraphID)
	if err != nil {
		return "", err
	}
	res, err := config.AlgoRPC.Pagerank(ctx, &algo.PageRankReq{Base: &algo.BaseReq{GraphID: req.GraphID, EdgeTags: edgeTags},
		Cfg: &algo.PRConfig{Iter: req.Iter, Prob: req.Prob}})
	if err != nil {
		return "", err
	}
	return HandleAlgoRPCRes(res, req.GraphID, taskID)
}
