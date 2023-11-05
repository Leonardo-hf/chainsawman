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

type CreateGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGraphLogic {
	return &CreateGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGraphLogic) CreateGraph(req *types.CreateGraphRequest) (resp *types.GraphInfoReply, err error) {
	resp = &types.GraphInfoReply{
		Graph: &types.Graph{},
	}
	idf := common.GraphCreate
	if req.TaskID != "" {
		// 任务已经提交过
		if err = util.FetchTask(l.ctx, l.svcCtx, req.TaskID, idf, resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
	graph := &model.Graph{
		Name:    req.Graph,
		GroupID: req.GroupID,
		Status:  common.GraphStatusInit,
	}
	err = l.svcCtx.MysqlClient.InsertGraph(l.ctx, graph)
	if err != nil {
		return nil, err
	}
	// 任务没提交过，创建任务
	req.GraphID = graph.ID
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, graph.ID, idf, req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.CreateGraph(req)
}
