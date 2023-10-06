package logic

import (
	"chainsawman/common"
	"chainsawman/graph/model"
	"context"
	"fmt"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlgoCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlgoCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlgoCreateLogic {
	return &AlgoCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlgoCreateLogic) AlgoCreate(req *types.CreateAlgoRequest) (resp *types.BaseReply, err error) {
	params := make([]*model.AlgosParam, 0)
	for _, p := range req.Algo.Params {
		params = append(params, &model.AlgosParam{
			FieldName: p.Key,
			FieldDesc: p.KeyDesc,
			FieldType: p.Type,
		})
	}
	err = l.svcCtx.MysqlClient.InsertAlgo(l.ctx, &model.Algo{
		Name:      req.Algo.Name,
		Desc:      req.Algo.Desc,
		Type:      req.Algo.Type,
		JarPath:   fmt.Sprintf("s3a://lib/%v", req.Jar),
		MainClass: req.EntryPoint,
		IsCustom:  common.Bool2Int64(true),
		Params:    params,
	})
	return nil, err
}
