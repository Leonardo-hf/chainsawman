package logic

import (
	"chainsawman/common"
	"context"
	"io"
	"sync"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHHILogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHHILogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHHILogic {
	return &GetHHILogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHHILogic) GetHHI() (resp *types.GetHHIReply, err error) {
	taskCnt := 1
	hhiCnt := 10
	tasks, err := l.svcCtx.MysqlClient.GetLatestHHITask(l.ctx, int64(taskCnt))
	if err != nil {
		return nil, err
	}
	wait := sync.WaitGroup{}
	hhiLanguage := make([]*types.HHILanguage, len(tasks))
	wait.Add(len(tasks))
	for i, t := range tasks {
		hhiLanguage[i] = &types.HHILanguage{
			Language:   t.Graph,
			UpdateTime: t.UpdateTime.UnixMilli(),
		}
		go func(filename string, i int) {
			hhi := make([]*types.HHI, 0)
			defer func() {
				hhiLanguage[i].HHIs = hhi
				wait.Done()
			}()
			file, err := l.svcCtx.OSSClient.FetchAlgo(l.ctx, filename)
			if err != nil {
				logx.Errorf("[Graph] fail to fetch algo file: %v", filename)
				return
			}
			// 解析文件
			parser, _ := common.NewExcelParser(file)
			for c := 0; c < hhiCnt; c++ {
				r, err := parser.Next()
				if err == io.EOF {
					break
				}
				class, _ := r.Get("class")
				score, _ := r.Get("score")
				hhi = append(hhi, &types.HHI{
					Name:  class,
					Score: score,
				})
			}
		}(t.Output, i)
	}
	wait.Wait()
	return &types.GetHHIReply{Languages: hhiLanguage}, nil
}
