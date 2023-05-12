package logic

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"chainsawman/graph/cmd/api/internal/util"
	"chainsawman/graph/model"

	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNeighborsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNeighborsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNeighborsLogic {
	return &GetNeighborsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNeighborsLogic) GetNeighbors(req *types.SearchNodeRequest) (resp *types.SearchNodeReply, err error) {
	resp = &types.SearchNodeReply{Base: &types.BaseReply{
		TaskID:     req.TaskID,
		TaskStatus: int64(model.KVTask_New),
	}}
	idf := common.GraphNeighbors
	if req.TaskID != 0 {
		// 任务已经提交过
		return resp, util.FetchTask(l.ctx, l.svcCtx, req.TaskID, idf, resp)
	}
	// 任务没提交过，创建任务
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, req.GraphID, idf, req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.GetNeighbors(req)
}
