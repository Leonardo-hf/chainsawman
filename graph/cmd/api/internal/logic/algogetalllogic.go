package logic

import (
	"chainsawman/graph/cmd/api/internal/types/rpc/algo"
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

func (l *AlgoGetAllLogic) AlgoGetAll() (resp *types.AlgoReply, err error) {
	algos, err := l.svcCtx.AlgoRPC.QueryAlgo(l.ctx, &algo.Empty{})
	if err != nil {
		return nil, err
	}
	sortAlgos := make(sortAlgo, 0)
	for _, a := range algos.GetAlgos() {
		params := make([]*types.Element, 0)
		for _, p := range a.GetParams() {
			params = append(params, &types.Element{
				Key:  p.GetKey(),
				Type: int64(p.GetType()),
			})
		}
		sortAlgos = append(sortAlgos, &types.Algo{
			Id:       a.GetId(),
			Name:     a.GetName(),
			Desc:     a.GetDesc(),
			IsCustom: a.GetIsCustom(),
			Params:   params,
		})
	}
	sort.Sort(sortAlgos)
	return &types.AlgoReply{
		Algos: sortAlgos,
	}, nil
}

type sortAlgo []*types.Algo

func (a sortAlgo) Len() int {
	return len(a)
}

func (a sortAlgo) Less(i int, j int) bool {
	return a[i].Type < a[j].Type || (!a[i].IsCustom && a[j].IsCustom) || a[i].Id < a[j].Id
}

func (a sortAlgo) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}
