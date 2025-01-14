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

func (l *GetGraphLogic) GetGraph(req *types.GetGraphDetailRequest) (resp *types.GetGraphDetailReply, err error) {
	resp = &types.GetGraphDetailReply{
		Base: &types.BaseReply{
			TaskID:     req.TaskID,
			TaskStatus: int64(model.KVTask_New),
		}}
	idf := common.GraphGet
	if req.TaskID != "" {
		// 任务已经提交过
		return resp, util.FetchTask(l.ctx, l.svcCtx, req.TaskID, resp)
	}
	// 任务没提交过，创建任务
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, req.GraphID, idf, req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.GetGraph(req)
}
