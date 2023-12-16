package logic

import (
	"chainsawman/common"
	"chainsawman/graph/model"
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupRequest) (resp *types.GroupInfoReply, err error) {
	group := &model.Group{
		Name:     req.Name,
		Desc:     req.Desc,
		ParentID: req.ParentID,
	}
	nodes := make([]*model.Node, 0)
	for _, n := range req.NodeTypeList {
		nodeAttrs := make([]*model.NodeAttr, 0)
		for _, a := range n.Attrs {
			nodeAttrs = append(nodeAttrs, &model.NodeAttr{
				Name: a.Name,
				Desc: a.Desc,
				Type: a.Type,
			})
		}
		nodes = append(nodes, &model.Node{
			Name:    n.Name,
			Desc:    n.Desc,
			Display: n.Display,
			Primary: n.Primary,
			Attrs:   nodeAttrs,
		})
	}
	group.Nodes = nodes
	edges := make([]*model.Edge, 0)
	for _, n := range req.EdgeTypeList {
		edgeAttrs := make([]*model.EdgeAttr, 0)
		for _, a := range n.Attrs {
			edgeAttrs = append(edgeAttrs, &model.EdgeAttr{
				Name: a.Name,
				Desc: a.Desc,
				Type: a.Type,
			})
		}
		edges = append(edges, &model.Edge{
			Name:    n.Name,
			Desc:    n.Desc,
			Display: n.Display,
			Direct:  common.Bool2Int64(n.EdgeDirection),
			Primary: n.Primary,
			Attrs:   edgeAttrs,
		})
	}
	group.Edges = edges
	err = l.svcCtx.MysqlClient.InsertGroup(l.ctx, group)
	return nil, err
}
