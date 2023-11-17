package db

import (
	"chainsawman/graph/db/query"
	"chainsawman/graph/model"
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlClientImpl struct {
}

type MysqlConfig struct {
	Addr string
}

func InitMysqlClient(cfg *MysqlConfig) MysqlClient {
	database, err := gorm.Open(mysql.Open(cfg.Addr))
	if err != nil {
		panic(err)
	}
	query.SetDefault(database)
	return &MysqlClientImpl{}
}

func (c *MysqlClientImpl) InsertTask(ctx context.Context, task *model.Task) error {
	t := query.Task
	task.Status = 0
	return t.WithContext(ctx).Select(t.Idf, t.Params, t.GraphID).Create(task)
}

func (c *MysqlClientImpl) UpdateTaskTIDByID(ctx context.Context, id int64, tid string) (int64, error) {
	t := query.Task
	info, err := t.WithContext(ctx).Where(t.ID.Eq(id)).Updates(map[string]interface{}{
		"tid": tid,
	})
	return info.RowsAffected, err
}

func (c *MysqlClientImpl) GetTasksByGraph(ctx context.Context, graphID int64) ([]*model.Task, error) {
	t := query.Task
	return t.WithContext(ctx).Where(t.GraphID.Eq(graphID)).Find()
}

func (c *MysqlClientImpl) GetTasks(ctx context.Context) ([]*model.Task, error) {
	t := query.Task
	return t.WithContext(ctx).Find()
}

func (c *MysqlClientImpl) GetTaskByID(ctx context.Context, id int64) (*model.Task, error) {
	t := query.Task
	return t.WithContext(ctx).Where(t.ID.Eq(id)).First()
}

func (c *MysqlClientImpl) InsertGraph(ctx context.Context, graph *model.Graph) error {
	g := query.Graph
	graph.Status = 0
	return g.WithContext(ctx).Create(graph)
}

func (c *MysqlClientImpl) DropGraphByID(ctx context.Context, id int64) (int64, error) {
	g := query.Graph
	res, err := g.WithContext(ctx).Where(g.ID.Eq(id)).Delete()
	return res.RowsAffected, err
}

func (c *MysqlClientImpl) DropTaskByID(ctx context.Context, id int64) (int64, error) {
	g := query.Task
	res, err := g.WithContext(ctx).Where(g.ID.Eq(id)).Delete()
	return res.RowsAffected, err
}

func (c *MysqlClientImpl) GetGraphByID(ctx context.Context, id int64) (*model.Graph, error) {
	g := query.Graph
	return g.WithContext(ctx).Select(g.NumNode, g.NumEdge).Where(g.ID.Eq(id)).First()
}

func (c *MysqlClientImpl) GetGraphByGroupID(ctx context.Context, groupID int64) ([]*model.Graph, error) {
	gr := query.Group
	g := query.Graph
	graphs, err := g.WithContext(ctx).Where(g.GroupID.Eq(groupID)).Find()
	if err != nil {
		return nil, err
	}
	children, err := gr.WithContext(ctx).Where(gr.ParentID.Eq(groupID)).Select(gr.ID).Find()
	if err != nil {
		return nil, err
	}
	for _, child := range children {
		childGraphs, err := c.GetGraphByGroupID(ctx, child.ID)
		if err != nil {
			return nil, err
		}
		graphs = append(graphs, childGraphs...)
	}
	return graphs, nil
}

func (c *MysqlClientImpl) GetGraphByName(ctx context.Context, name string) (*model.Graph, error) {
	g := query.Graph
	return g.WithContext(ctx).Where(g.Name.Eq(name)).First()
}

func (c *MysqlClientImpl) UpdateGraphStatusByID(ctx context.Context, id int64, status int64) (int64, error) {
	g := query.Graph
	res, err := g.WithContext(ctx).Where(g.ID.Eq(id)).Updates(map[string]interface{}{
		"status": status,
	})
	return res.RowsAffected, err
}

func (c *MysqlClientImpl) GetAllGraph(ctx context.Context) ([]*model.Graph, error) {
	g := query.Graph
	return g.WithContext(ctx).Find()
}

func (c *MysqlClientImpl) InsertGroup(ctx context.Context, group *model.Group) error {
	gr := query.Group
	return gr.WithContext(ctx).Create(group)
}

func (c *MysqlClientImpl) DropGroupByID(ctx context.Context, id int64) (int64, error) {
	gr := query.Group
	res, err := gr.WithContext(ctx).Where(gr.ID.Eq(id)).Delete()
	return res.RowsAffected, err
}

func (c *MysqlClientImpl) GetAllGroups(ctx context.Context) ([]*model.Group, error) {
	gr := query.Group
	return gr.WithContext(ctx).Preload(gr.Nodes.NodeAttrs).Preload(gr.Edges.EdgeAttrs).Find()
}

func (c *MysqlClientImpl) GetNodeByID(ctx context.Context, id int64) (*model.Node, error) {
	n := query.Node
	return n.WithContext(ctx).Where(n.ID.Eq(id)).Preload(n.NodeAttrs).First()
}

func (c *MysqlClientImpl) GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error) {
	g := query.Graph
	gr := query.Group
	sub := g.WithContext(ctx).Select(g.GroupID).Where(g.ID.Eq(id))
	group, err := gr.WithContext(ctx).Preload(gr.Nodes.NodeAttrs).Preload(gr.Edges.EdgeAttrs).Where(gr.Columns(gr.ID).Eq(sub)).First()
	return group, err
}

func (c *MysqlClientImpl) GetAllAlgo(ctx context.Context) ([]*model.Algo, error) {
	a := query.Algo
	return a.WithContext(ctx).Select(a.ID, a.Desc, a.IsCustom, a.GroupID, a.Name, a.Type).Preload(a.Params).Find()
}

func (c *MysqlClientImpl) DropAlgoByID(ctx context.Context, id int64) (int64, error) {
	a := query.Algo
	res, err := a.WithContext(ctx).Where(a.ID.Eq(id)).Delete()
	return res.RowsAffected, err
}

func (c *MysqlClientImpl) InsertAlgo(ctx context.Context, algo *model.Algo) error {
	a := query.Algo
	return a.WithContext(ctx).Create(algo)
}
