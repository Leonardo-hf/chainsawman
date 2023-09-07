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

type UpdateGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGraphLogic {
	return &UpdateGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGraphLogic) UpdateGraph(req *types.UpdateGraphRequest) (resp *types.GraphInfoReply, err error) {
	resp = &types.GraphInfoReply{
		Base: &types.BaseReply{
			TaskID:     req.TaskID,
			TaskStatus: int64(model.KVTask_New),
		},
	}
	idf := common.GraphUpdate
	if req.TaskID != "" {
		// 任务已经提交过
		if err = util.FetchTask(l.ctx, l.svcCtx, req.TaskID, idf, resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
	_, err = l.svcCtx.MysqlClient.UpdateGraphStatusByID(l.ctx, req.GraphID, common.GraphStatusUpdate)
	if err != nil {
		return nil, err
	}
	// 任务没提交过，创建任务
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, req.GraphID, idf, req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.UpdateGraph(req)
}
