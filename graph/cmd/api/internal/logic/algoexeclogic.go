package logic

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/config"
	"chainsawman/graph/cmd/api/internal/util"
	"chainsawman/graph/model"
	"context"
	"strconv"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlgoExecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlgoExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlgoExecLogic {
	return &AlgoExecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlgoExecLogic) AlgoExec(req *types.ExecAlgoRequest) (resp *types.AlgoReply, err error) {
	// 参数校验
	for _, p := range req.Params {
		switch p.Type {
		case common.TypeString:
			if len(p.Value) == 0 {
				return nil, config.ErrInvalidParam(p.Key, "blank string")
			}
			break
		case common.TypeInt, common.TypeDouble:
			_, err := strconv.ParseFloat(p.Value, 64)
			if err != nil {
				return nil, config.ErrInvalidParam(p.Key, "should be int or double")
			}
			break
		case common.TypeStringList:
			for _, pi := range p.ListValue {
				if len(pi) == 0 {
					return nil, config.ErrInvalidParam(p.Key, "blank string in list")
				}
			}
			break
		case common.TypeDoubleList:
			for _, pi := range p.ListValue {
				_, err := strconv.ParseFloat(pi, 64)
				if err != nil {
					return nil, config.ErrInvalidParam(p.Key, "should be int or double list")
				}
			}
		}
	}
	// 设置返回值
	resp = &types.AlgoReply{Base: &types.BaseReply{
		TaskID:     req.TaskID,
		TaskStatus: int64(model.KVTask_New),
	}}
	idf := common.AlgoExec
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
	return l.AlgoExec(req)
}
