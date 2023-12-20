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

// TODO: 什么硬编码？
func (c *MysqlClientImpl) GetLatestSETask(ctx context.Context, n int64) ([]*SETask, error) {
	e := query.Exec
	a := query.Algo
	g := query.Graph
	res := make([]*SETask, 0)
	err := e.WithContext(ctx).Join(a, a.ID.EqCol(e.AlgoID)).Join(g, g.ID.EqCol(g.ID)).Distinct(e.AlgoID, e.GraphID).
		Select(e.Output, e.UpdateTime, e.GraphID, g.Name.As("graph"), a.Name.As("algo")).Where(e.Status.Eq(1)).Where(a.Tag.Eq("软件影响力")).
		Order(e.UpdateTime.Desc()).Limit(int(n)).Scan(&res)
	return res, err
}

// TODO: 什么硬编码？
func (c *MysqlClientImpl) GetLatestHHITask(ctx context.Context, n int64) ([]*HHITask, error) {
	e := query.Exec
	a := query.Algo
	g := query.Graph
	res := make([]*HHITask, 0)
	sub := a.WithContext(ctx).Select(a.ID).Where(a.Name.Eq("hhi"))
	err := e.WithContext(ctx).Join(g, g.ID.EqCol(e.GraphID)).Distinct(e.GraphID).
		Select(e.Output, e.UpdateTime, g.Name.As("graph")).Where(e.Status.Eq(1)).Where(e.Columns(e.AlgoID).In(sub)).
		Order(e.UpdateTime.Desc()).Limit(int(n)).Scan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MysqlClientImpl) InsertAlgoTask(ctx context.Context, exec *model.Exec) error {
	e := query.Exec
	return e.WithContext(ctx).Create(exec)
}

func (c *MysqlClientImpl) GetAlgoTaskByID(ctx context.Context, id int64) (*model.Exec, error) {
	t := query.Exec
	return t.WithContext(ctx).Where(t.ID.Eq(id)).First()
}

func (c *MysqlClientImpl) DropAlgoTaskByID(ctx context.Context, id int64) (int64, error) {
	g := query.Exec
	res, err := g.WithContext(ctx).Where(g.ID.Eq(id)).Delete()
	return res.RowsAffected, err
}

func (c *MysqlClientImpl) GetAlgoTasks(ctx context.Context) ([]*model.Exec, error) {
	t := query.Exec
	return t.WithContext(ctx).Find()
}

func (c *MysqlClientImpl) GetAlgoTasksByGraphId(ctx context.Context, id int64) ([]*model.Exec, error) {
	t := query.Exec
	return t.WithContext(ctx).Where(t.GraphID.Eq(id)).Find()
}

func (c *MysqlClientImpl) UpdateAlgoTaskTIDByID(ctx context.Context, id int64, tid string) (int64, error) {
	t := query.Exec
	info, err := t.WithContext(ctx).Where(t.ID.Eq(id)).Updates(map[string]interface{}{
		"tid": tid,
	})
	return info.RowsAffected, err
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
	ids := make([]int64, 0)
	q := []int64{id}
	for len(q) > 0 {
		ids = append(ids, q...)
		children, err := gr.WithContext(ctx).Select(gr.ID).Where(gr.ParentID.In(q...)).Find()
		if err != nil {
			return 0, err
		}
		// 清空 q
		q = q[:0]
		for _, child := range children {
			q = append(q, child.ID)
		}
	}
	res, err := gr.WithContext(ctx).Where(gr.ID.In(ids...)).Delete()
	return res.RowsAffected, err
}

func (c *MysqlClientImpl) GetAllGroups(ctx context.Context) ([]*model.Group, error) {
	gr := query.Group
	return gr.WithContext(ctx).Preload(gr.Nodes.Attrs).Preload(gr.Edges.Attrs).Find()
}

func (c *MysqlClientImpl) GetNodeByID(ctx context.Context, id int64) (*model.Node, error) {
	n := query.Node
	return n.WithContext(ctx).Where(n.ID.Eq(id)).Preload(n.Attrs).First()
}

func (c *MysqlClientImpl) GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error) {
	g := query.Graph
	gr := query.Group
	sub := g.WithContext(ctx).Select(g.GroupID).Where(g.ID.Eq(id))
	group, err := gr.WithContext(ctx).Preload(gr.Nodes.Attrs).Preload(gr.Edges.Attrs).Where(gr.Columns(gr.ID).Eq(sub)).First()
	return group, err
}

func (c *MysqlClientImpl) GetAllAlgo(ctx context.Context) ([]*model.Algo, error) {
	a := query.Algo
	return a.WithContext(ctx).Select(a.ID, a.GroupID, a.Name, a.Tag, a.Detail).Preload(a.Params).Find()
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

func (c *MysqlClientImpl) GetAlgoExecCfgByID(ctx context.Context, id int64) (*model.Algo, error) {
	a := query.Algo
	return a.WithContext(ctx).Select(a.MainClass, a.JarPath).Where(a.ID.Eq(id)).First()
}
