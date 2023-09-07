package handler

import (
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"chainsawman/consumer/task/types/rpc/algo"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoBetweenness struct {
}

func (h *AlgoBetweenness) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.AlgoRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	edgeTags, err := PrepareForAlgoRPC(ctx, req.GraphID)
	if err != nil {
		return "", err
	}
	res, err := config.AlgoRPC.Betweenness(ctx, &algo.BaseReq{GraphID: req.GraphID, EdgeTags: edgeTags})
	if err != nil {
		return "", err
	}
	return HandleAlgoRPCRes(res, req.GraphID, taskID)
}

func PrepareForAlgoRPC(ctx context.Context, graphID int64) ([]string, error) {
	group, err := config.MysqlClient.GetGroupByGraphId(ctx, graphID)
	if err != nil {
		return nil, err
	}
	edgeTags := make([]string, len(group.Edges))
	for i, e := range group.Edges {
		edgeTags[i] = e.Name
	}
	return edgeTags, nil
}

func HandleAlgoRPCRes(res *algo.RankReply, graphID int64, taskID string) (string, error) {
	ids := make([]int64, len(res.GetRanks()))
	ranks := make([]*types.Rank, len(res.GetRanks()))
	for i, r := range res.GetRanks() {
		ids[i] = r.Id
		ranks[i] = &types.Rank{Score: r.Score}
	}
	nodePacks, err := config.NebulaClient.GetNodesByIds(graphID, ids)
	if err != nil {
		return "", err
	}
	// 一个垃圾的遍历赋值
	for t, np := range nodePacks {
		for _, n := range np {
			for i, id := range ids {
				if id == n.Id {
					ranks[i].Node = n
					ranks[i].Tag = t
					break
				}
			}
		}
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
