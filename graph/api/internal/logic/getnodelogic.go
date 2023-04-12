package logic

import (
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"
	"chainsawman/graph/api/internal/util"
	"chainsawman/graph/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeLogic {
	return &GetNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeLogic) GetNode(req *types.SearchNodeRequest) (resp *types.SearchNodeReply, err error) {
	resp = &types.SearchNodeReply{Base: &types.BaseReply{
		TaskID:     req.TaskID,
		TaskStatus: int64(model.KVTask_New),
	}}
	if req.TaskID != 0 {
		// 任务已经提交过
		return resp, util.FetchTask(l.ctx, l.svcCtx, req.TaskID, resp)
	}
	// 任务没提交过，创建任务
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, "GetNode", req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.GetNode(req)
}
