package handler

import (
	"chainsawman/consumer/config"
	"chainsawman/consumer/model"
	"chainsawman/consumer/types"
	"chainsawman/consumer/types/rpc/algo"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type AlgoAvgCC struct {
}

func (h *AlgoAvgCC) Handle(params string, taskID int64) (string, error) {
	req := &types.AlgoRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	res, err := config.AlgoRPC.AvgClustering(ctx, &algo.BaseReq{GraphID: req.GraphID})
	if err != nil {
		return "", err
	}
	resp := &types.AlgoMetricReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Score: res.Score,
	}
	return jsonx.MarshalToString(resp)
}
