package handler

import (
	"chainsawman/consumer/config"
	"chainsawman/consumer/model"
	"chainsawman/consumer/types"
	"chainsawman/consumer/types/rpc/algo"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoPageRank struct {
}

func (h *AlgoPageRank) Handle(params string, taskID int64) (string, error) {
	req := &types.AlgoDegreeRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	// TODO: 从参数中获取cfg
	res, err := config.AlgoRPC.Pagerank(ctx, &algo.PageRankReq{Base: &algo.BaseReq{GraphID: req.GraphID}, Cfg: &algo.PRConfig{Iter: 3, Prob: 0.85}})
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
	}
	return jsonx.MarshalToString(resp)
}
