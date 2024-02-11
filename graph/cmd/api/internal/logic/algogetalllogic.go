package logic

import (
	"context"
	"sort"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlgoGetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlgoGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlgoGetAllLogic {
	return &AlgoGetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlgoGetAllLogic) AlgoGetAll() (resp *types.GetAlgoReply, err error) {
	algos, err := l.svcCtx.MysqlClient.GetAllAlgo(l.ctx)
	if err != nil {
		return nil, err
	}
	sortAlgos := make(sortAlgo, 0)
	for _, a := range algos {
		params := make([]*types.AlgoParam, 0)
		for _, p := range a.Params {
			params = append(params, &types.AlgoParam{
				Key:       p.Name,
				KeyDesc:   p.Desc,
				Type:      p.Type,
				InitValue: p.Default,
				Max:       p.Max,
				Min:       p.Min,
			})
		}
		sortAlgos = append(sortAlgos, &types.Algo{
			Id:      a.ID,
			Name:    a.Name,
			Define:  a.Define,
			GroupId: a.GroupID,
			Params:  params,
			Tag:     a.Tag,
			TagID:   a.TagID,
		})
	}
	sort.Sort(sortAlgos)
	return &types.GetAlgoReply{
		Algos: sortAlgos,
	}, nil
}

type sortAlgo []*types.Algo

func (a sortAlgo) Len() int {
	return len(a)
}

func (a sortAlgo) Less(i int, j int) bool {
	return a[i].Tag < a[j].Tag || a[i].Id < a[j].Id
}

func (a sortAlgo) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}
