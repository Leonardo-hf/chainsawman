package logic

import (
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"
	"chainsawman/graph/config"
	"chainsawman/graph/model"

	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeLogic {
	return &GetNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeLogic) GetNode(req *types.SearchNodeRequest) (resp *types.SearchNodeReply, err error) {
	// TODO: 这个方法里的先查redis再查mysql之类的流程可以抽象成通用的方法
	resp = &types.SearchNodeReply{TaskID: req.TaskID}
	if req.TaskID != 0 {
		// 任务已经提交过
		var result string
		status := 0
		task, err := config.RedisClient.GetTaskById(l.ctx, req.TaskID)
		if err == nil {
			// redis里有记录
			result = task.Result
			status = int(task.Status)
		} else if err == redis.Nil {
			// redis里没有记录，查询mysql
			oTask, err := config.MysqlClient.SearchTaskById(req.TaskID)
			if err != nil {
				return nil, err
			}
			status = int(oTask.Status)
			result = oTask.Result
			// 更新redis
			err = config.RedisClient.UpsertTask(l.ctx, &model.KVTask{
				Id:         oTask.ID,
				Name:       oTask.Name,
				Params:     oTask.Params,
				Status:     oTask.Status,
				CreateTime: oTask.CreateTime,
				UpdateTime: oTask.UpdateTime,
				Result:     oTask.Result,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
		// 任务完成了
		if status == 1 {
			err = jsonx.UnmarshalFromString(result, resp)
			if err != nil {
				return nil, err
			}
		}
		return resp, nil
	}
	// 任务没提交过，创建任务
	params, err := jsonx.MarshalToString(req)
	if err != nil {
		return nil, err
	}
	oTask := &model.Task{
		Params: params,
		Name:   "GetNode",
	}
	err = config.MysqlClient.InsertTask(oTask)
	if err != nil {
		return nil, err
	}
	// 发布任务
	err = config.RedisClient.ProduceTaskMsg(l.ctx, &model.KVTask{
		Id:         oTask.ID,
		Name:       oTask.Name,
		Params:     oTask.Params,
		Status:     oTask.Status,
		CreateTime: oTask.CreateTime,
		UpdateTime: oTask.UpdateTime,
	})
	if err != nil {
		return nil, err
	}
	// 重试一次
	req.TaskID = oTask.ID
	return l.GetNode(req)
}
