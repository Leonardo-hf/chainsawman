package logic

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/config"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"chainsawman/graph/model"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"strconv"

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

func (l *AlgoExecLogic) AlgoExec(req *types.ExecAlgoRequest) (resp *types.GetAlgoTaskReply, err error) {
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
	// algoID 查找文件路径
	execCfg, err := l.svcCtx.MysqlClient.GetAlgoExecCfgByID(l.ctx, req.AlgoID)
	if err != nil {
		return nil, err
	}
	// 生成文件名称
	fileName := fmt.Sprintf("%v-%v-%v", req.GraphID, req.AlgoID, uuid.New().String())
	execParams := params2Map(req.Params)
	execParams["graphID"] = fmt.Sprintf("G%v", req.GraphID)
	execParams["target"] = fileName
	// 提交任务
	appID, err := l.svcCtx.AlgoService.SubmitAlgo(execCfg.JarPath, execCfg.MainClass, execParams)
	if err != nil {
		return nil, err
	}
	logx.Infof("[Graph] start algo job, appID: %v", appID)
	params, _ := jsonx.MarshalToString(req.Params)
	exec := &model.Exec{
		AlgoID:  req.AlgoID,
		Status:  int64(model.KVTask_New),
		Params:  params,
		GraphID: req.GraphID,
		Output:  fileName,
		AppID:   appID,
	}
	err = l.svcCtx.MysqlClient.InsertAlgoTask(l.ctx, exec)
	return &types.GetAlgoTaskReply{
		Task: &types.AlgoTask{
			Id:         exec.ID,
			CreateTime: exec.CreateTime.UnixMilli(),
			UpdateTime: exec.UpdateTime.UnixMilli(),
			GraphID:    req.GraphID,
			Req:        params,
			AlgoID:     req.AlgoID,
		},
	}, err
}

func params2Map(params []*types.Param) map[string]interface{} {
	m := make(map[string]interface{})
	for _, p := range params {
		switch p.Type {
		case common.TypeString, common.TypeInt, common.TypeDouble:
			m[p.Key] = p.Value
			break
		case common.TypeDoubleList, common.TypeStringList:
			m[p.Key] = p.ListValue
			break
		}
	}
	return m
}
