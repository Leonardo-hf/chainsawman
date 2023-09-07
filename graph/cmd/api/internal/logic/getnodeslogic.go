package logic

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/util"
	"chainsawman/graph/model"
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodesLogic {
	return &GetNodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodesLogic) GetNodes(req *types.GetNodesRequest) (resp *types.GetNodesReply, err error) {
	resp = &types.GetNodesReply{
		Base: &types.BaseReply{
			TaskID:     req.TaskID,
			TaskStatus: int64(model.KVTask_New),
		},
	}
	idf := common.GraphNodes
	if req.TaskID != "" {
		// 任务已经提交过
		if err = util.FetchTask(l.ctx, l.svcCtx, req.TaskID, idf, resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
	// 任务没提交过，创建任务
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, req.GraphID, idf, req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.GetNodes(req)
}
