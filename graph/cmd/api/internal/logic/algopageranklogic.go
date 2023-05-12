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

type AlgoPageRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlgoPageRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlgoPageRankLogic {
	return &AlgoPageRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlgoPageRankLogic) AlgoPageRank(req *types.AlgoPageRankRequest) (resp *types.AlgoRankReply, err error) {
	resp = &types.AlgoRankReply{Base: &types.BaseReply{
		TaskID:     req.TaskID,
		TaskStatus: int64(model.KVTask_New),
	}}
	idf := common.AlgoPagerank
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
	return l.AlgoPageRank(req)
}
