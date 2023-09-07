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
	edgeTags, err := PrepareForAlgoRPC(ctx, req.GraphID)
	if err != nil {
		return "", err
	}
	res, err := config.AlgoRPC.Louvain(ctx, &algo.LouvainReq{
		Base: &algo.BaseReq{GraphID: req.GraphID, EdgeTags: edgeTags},
		Cfg: &algo.LouvainConfig{
			MaxIter:      req.MaxIter,
			InternalIter: req.InternalIter,
			Tol:          req.Tol,
		},
	})
	if err != nil {
		return "", err
	}
	return HandleAlgoRPCRes(res, req.GraphID, taskID)
}
