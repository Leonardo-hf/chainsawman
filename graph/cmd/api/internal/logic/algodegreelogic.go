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

type AlgoDegreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlgoDegreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlgoDegreeLogic {
	return &AlgoDegreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlgoDegreeLogic) AlgoDegree(req *types.AlgoDegreeRequest) (resp *types.AlgoRankReply, err error) {
	resp = &types.AlgoRankReply{Base: &types.BaseReply{
		TaskID:     req.TaskID,
		TaskStatus: int64(model.KVTask_New),
	}}
	idf := common.AlgoDegree
	if req.TaskID != "" {
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
	return l.AlgoDegree(req)
}
