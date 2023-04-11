package logic

import (
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"
	"chainsawman/graph/api/internal/util"

	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGraphLogic {
	return &GetGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGraphLogic) GetGraph(req *types.SearchRequest) (resp *types.SearchGraphDetailReply, err error) {
	resp = &types.SearchGraphDetailReply{}
	if req.TaskID != 0 {
		// 任务已经提交过
		return resp, util.FetchTask(l.ctx, l.svcCtx, req.TaskID, resp)
	}
	// 任务没提交过，创建任务
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, "GetGraph", req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.GetGraph(req)
}
