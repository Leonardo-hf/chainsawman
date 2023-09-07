package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type CreateGraph struct {
}

func (h *CreateGraph) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.CreateGraphRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	group, err := config.MysqlClient.GetGroupByID(ctx, req.GroupID)
	// 创建图空间
	err = config.NebulaClient.CreateGraph(req.GraphID, group)
	if err != nil {
		return "", err
	}
	// 尝试三次
	ok := false
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Second)
		ok, err = config.NebulaClient.HasGraph(req.GraphID)
		if err != nil {
			logx.Error(err)
		}
		if ok {
			break
		}
	}
	_, err = config.MysqlClient.UpdateGraphByID(ctx, &model.Graph{
		ID:      req.GraphID,
		Status:  common.GraphStatusOK,
		NumNode: 0,
		NumEdge: 0,
	})
	if err != nil {
		logx.Error(err)
	}
	return jsonx.MarshalToString(types.GraphInfoReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Graph: &types.Graph{
			Name:    req.Graph,
			Desc:    req.Desc,
			Id:      req.GraphID,
			GroupID: req.GroupID,
		}})
}
