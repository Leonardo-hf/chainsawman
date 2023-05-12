package handler

import (
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"chainsawman/consumer/task/types/rpc/algo"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoVoteRank struct {
}

func (h *AlgoVoteRank) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.AlgoVoteRankRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	res, err := config.AlgoRPC.Voterank(ctx, &algo.VoteRankReq{
		Base: &algo.BaseReq{GraphID: req.GraphID},
		Cfg: &algo.VoteConfig{
			Iter: req.Iter,
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
