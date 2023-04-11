package logic

import (
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"
	"chainsawman/graph/api/internal/util"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 1 << 31

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest) (resp *types.SearchGraphReply, err error) {
	resp = &types.SearchGraphReply{}
	if req.TaskID != 0 {
		// 任务已经提交过
		return resp, util.FetchTask(l.ctx, l.svcCtx, req.TaskID, resp)
	}
	// 任务没提交过，创建任务
	taskID, err := util.PublishTask(l.ctx, l.svcCtx, "Upload", req)
	if err != nil {
		return nil, err
	}
	req.TaskID = taskID
	// 重试一次
	return l.Upload(req)
}
