package logic

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllGraphLogic {
	return &GetAllGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllGraphLogic) GetAllGraph() (resp *types.GetAllGraphReply, err error) {
	groups, err := l.svcCtx.MysqlClient.GetAllGroups(l.ctx)
	if err != nil {
		return nil, err
	}
	graphs, err := l.svcCtx.MysqlClient.GetAllGraph(l.ctx)
	if err != nil {
		return nil, err
	}
	gmap := make(map[int64]*types.Group)
	resp = &types.GetAllGraphReply{
		Groups: make([]*types.Group, 0),
	}
	// 获得所有group
	for _, group := range groups {
		nodes, edges := make([]*types.Structure, 0), make([]*types.Structure, 0)
		for _, n := range group.Nodes {
			nodeAttrs := make([]*types.Attr, 0)
			for _, a := range n.Attrs {
				nodeAttrs = append(nodeAttrs, &types.Attr{
					Name: a.Name,
					Desc: a.Desc,
					Type: a.Type,
				})
			}
			nodes = append(nodes, &types.Structure{
				Id:      n.ID,
				Name:    n.Name,
				Desc:    n.Desc,
				Display: n.Display,
				Attrs:   nodeAttrs,
			})
		}
		for _, n := range group.Edges {
			edgeAttrs := make([]*types.Attr, 0)
			for _, a := range n.Attrs {
				edgeAttrs = append(edgeAttrs, &types.Attr{
					Name: a.Name,
					Desc: a.Desc,
					Type: a.Type,
				})
			}
			edges = append(edges, &types.Structure{
				Id:            n.ID,
				Name:          n.Name,
				Desc:          n.Desc,
				Display:       n.Display,
				EdgeDirection: common.Int642Bool(n.Direct),
				Attrs:         edgeAttrs,
			})
		}
		gmap[group.ID] = &types.Group{
			Id:           group.ID,
			Name:         group.Name,
			Desc:         group.Desc,
			ParentID:     group.ParentID,
			NodeTypeList: nodes,
			EdgeTypeList: edges,
			Graphs:       make([]*types.Graph, 0),
		}
	}
	// 为所有group填充graph
	for _, g := range graphs {
		gmap[g.GroupID].Graphs = append(gmap[g.GroupID].Graphs, &types.Graph{
			Id:       g.ID,
			Status:   g.Status,
			GroupID:  g.GroupID,
			Name:     g.Name,
			NumNode:  g.NumNode,
			NumEdge:  g.NumEdge,
			CreatAt:  g.CreateTime.UnixMilli(),
			UpdateAt: g.UpdateTime.UnixMilli(),
		})
	}
	// map -> list
	for _, v := range gmap {
		resp.Groups = append(resp.Groups, v)
	}
	return resp, nil
}
