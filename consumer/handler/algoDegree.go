package handler

import (
	"chainsawman/consumer/config"
	"chainsawman/consumer/model"
	"chainsawman/consumer/types"
	"chainsawman/consumer/types/rpc/algo"

	"context"

	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoDegree struct {
}

func (h *AlgoDegree) Handle(params string, taskID int64) (string, error) {
	req := &types.AlgoDegreeRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	res, err := config.AlgoRPC.Degree(ctx, &algo.BaseReq{GraphID: req.GraphID})
	if err != nil {
		return "", err
	}

	ranks := make([]*types.Rank, 0)
	for _, r := range res.GetRanks() {
		ranks = append(ranks, &types.Rank{
			Node:  r.Name,
			Score: r.Score,
		})
	}
	resp := &types.AlgoRankReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Ranks: ranks,
	}
	return jsonx.MarshalToString(resp)
}
