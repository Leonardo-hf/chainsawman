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

type GetHotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotLogic {
	return &GetHotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotLogic) GetHot() (resp *types.GetHotSEReply, err error) {
	taskCnt := 3
	seCnt := 10
	tasks, err := l.svcCtx.MysqlClient.GetLatestSETask(l.ctx, int64(taskCnt))
	if err != nil {
		return nil, err
	}
	wait := sync.WaitGroup{}
	hots := make([]*types.HotSETopic, len(tasks))
	wait.Add(len(tasks))
	for i, t := range tasks {
		hots[i] = &types.HotSETopic{
			Language:   t.Graph,
			Topic:      t.Algo,
			UpdateTime: t.UpdateTime.UnixMilli(),
		}
		go func(graph int64, filename string, i int) {
			se := make([]*types.HotSE, 0)
			defer func() {
				hots[i].Software = se
				wait.Done()
			}()
			rids := make([]int64, 0)
			file, err := l.svcCtx.OSSClient.FetchAlgo(l.ctx, filename)
			if err != nil {
				logx.Errorf("[Graph] fail to fetch algo file: %v", filename)
				return
			}
			// 解析文件
			parser, _ := common.NewExcelParser(file)
			for c := 0; c < seCnt; c++ {
				r, err := parser.Next()
				if err == io.EOF {
					break
				}
				artifact, _ := r.Get("artifact")
				version, _ := r.Get("version")
				score, _ := r.GetAsFloat64("score")
				se = append(se, &types.HotSE{
					Artifact: artifact,
					Version:  version,
					Score:    score,
				})
			}
			ds, err := l.svcCtx.NebulaClient.GetLibraryByReleaseIDs(graph, rids)
			if err != nil {
				logx.Errorf("[Graph] fail to get library by rids: %v", rids)
				return
			}
			for p, rid := range rids {
				if library, ok := ds[rid]; ok {
					se[p].HomePage = library.Homepage
				}
			}
		}(t.GraphID, t.Output, i)
	}
	wait.Wait()
	return &types.GetHotSEReply{Topics: hots}, nil
}
