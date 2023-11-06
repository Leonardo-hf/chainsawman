package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"context"
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"strconv"
	"time"
)

type AlgoExec struct {
}

func (h *AlgoExec) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.ExecAlgoRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	// algoID 查找文件路径
	execCfg, err := config.MysqlClient.GetAlgoExecCfgByID(ctx, req.AlgoID)
	if err != nil {
		return "", err
	}
	// 生成文件名称
	fileName := fmt.Sprintf("%v-%v-%v", req.GraphID, req.AlgoID, uuid.New().String())
	// 查 edgeTags
	group, err := config.MysqlClient.GetGroupByGraphId(ctx, req.GraphID)
	if err != nil {
		return "", err
	}
	edgeTags := make([]string, len(group.Edges))
	for i, e := range group.Edges {
		edgeTags[i] = e.Name
	}
	execParams := params2Map(req.Params)
	execParams["graphID"] = fmt.Sprintf("G%v", req.GraphID)
	execParams["edgeTags"] = edgeTags
	execParams["target"] = fileName
	// 提交任务，要防止重复提交任务
	appID, err := config.AlgoService.SubmitAlgo(execCfg.JarPath, execCfg.MainClass, execParams)
	if err != nil {
		return "", err
	}
	//  轮询任务是否完成
	go checkIsFinish(ctx, task, appID, fileName, req.GraphID, func(graphID int64, output string) string {
		res, _ := jsonx.MarshalToString(&types.AlgoReply{
			Base: &types.BaseReply{
				TaskID:     taskID,
				TaskStatus: int64(model.KVTask_Finished),
			},
			File:   fileName,
			AlgoID: req.AlgoID,
		})
		return res
	})
	return "", config.DelayTaskErr
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
		case common.TypeRankAlgo, common.TypeScoreAlgo:
			v := make([]map[string]string, 0)
			content, err := config.OSSClient.FetchAlgo(context.Background(), p.AlgoValue)
			if err != nil {
				logx.ErrorStack(err)
			}
			parser, err := common.NewExcelParser(content)
			if err != nil {
				logx.ErrorStack(err)
			}
			for {
				r, err := parser.Next()
				if err == io.EOF {
					break
				}
				vi := make(map[string]string)
				for _, f := range r.Keys() {
					s, _ := r.Get(f)
					vi[f] = s
				}
				v = append(v, vi)
			}
			m[p.Key] = v
		}
	}
	return m
}

func checkIsFinish(ctx context.Context, task *model.KVTask, appID string, output string, graphID int64,
	ret func(graphID int64, output string) string) {
	op := func() error {
		// 判断任务是否还在，还在就继续
		taskIDInt, _ := strconv.ParseInt(task.Id, 10, 64)
		has, err := config.MysqlClient.HasTaskByID(ctx, taskIDInt)
		if err != nil {
			logx.ErrorStack(err)
		}
		// TODO: 任务被删除了，这里也停止spark
		if !has {
			err = config.AlgoService.StopAlgo(appID)
			if err != nil {
				logx.ErrorStack(err)
			}
			return nil
		}
		ok := config.OSSClient.AlgoGenerated(ctx, output)
		if ok {
			res := ret(graphID, output)
			task.Result = res
			task.Status = model.KVTask_Finished
			task.UpdateTime = time.Now().UTC().Unix()
			if err = config.RedisClient.UpsertTask(ctx, task); err != nil {
				return err
			}
			if common.TaskIdf(task.Idf).Persistent {
				taskIDInt, _ := strconv.ParseInt(task.Id, 10, 64)
				_, err = config.MysqlClient.UpdateTaskByID(ctx, &model.Task{
					ID:     taskIDInt,
					Status: int64(task.Status),
					Result: res,
				})
			}
			return nil
		}
		return errors.New("not ready")
	}
	backoffCfg := backoff.NewExponentialBackOff()
	backoffCfg.InitialInterval = time.Minute
	backoffCfg.MaxElapsedTime = 365 * 24 * time.Hour
	backoffCfg.MaxInterval = 10 * time.Minute
	_ = backoff.Retry(op, backoffCfg)
	logx.Infof("[Task] finish task, idf=%v", task.Idf)
}
