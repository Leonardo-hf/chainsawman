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

func (l *CreateGraphLogic) CreateGraph(req *types.UploadRequest) (resp *types.SearchGraphReply, err error) {
	resp = &types.SearchGraphReply{
		Base: &types.BaseReply{
			TaskID:     req.TaskID,
			TaskStatus: int64(model.KVTask_New),
		},
		Graph: &types.Graph{},
	}
	if req.TaskID != 0 {
		// 任务已经提交过
		if err = util.FetchTask(l.ctx, l.svcCtx, req.TaskID, resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
	graph := &model.Graph{
		Name: req.Graph,
		Desc: req.Desc,
	}
	err = l.svcCtx.MysqlClient.InsertGraph(l.ctx, graph)
	if err != nil {
		return nil, err
	}
	// 任务没提交过，创建任务
	req.GraphID = graph.ID
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, graph.ID, common.GraphCreate, req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.CreateGraph(req)
}
