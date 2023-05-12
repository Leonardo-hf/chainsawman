package handler

import (
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"chainsawman/consumer/task/types/rpc/algo"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoLouvain struct {
}

func (h *AlgoLouvain) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.AlgoLouvainRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	res, err := config.AlgoRPC.Louvain(ctx, &algo.LouvainReq{
		Base: &algo.BaseReq{GraphID: req.GraphID},
		Cfg: &algo.LouvainConfig{
			MaxIter:      req.MaxIter,
			InternalIter: req.InternalIter,
			Tol:          req.Tol,
		},
	})
	if err != nil {
		return "", err
	}
	ranks := make([]*types.Rank, 0)
	for _, r := range res.GetRanks() {
		ranks = append(ranks, &types.Rank{
			NodeID: r.Id,
			Score:  r.Score,
		})
	}
	resp := &types.AlgoRankReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Ranks: ranks,
		File:  res.GetFile(),
	}
	return jsonx.MarshalToString(resp)
}
