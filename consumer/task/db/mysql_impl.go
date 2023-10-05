package db

import (
	"chainsawman/consumer/task/db/query"
	"chainsawman/consumer/task/model"
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

func (c *MysqlClientImpl) HasTaskByID(ctx context.Context, id int64) (bool, error) {
	t := query.Task
	cnt, err := t.WithContext(ctx).Where(t.ID.Eq(id)).Count()
	return cnt == 1, err
}

func (c *MysqlClientImpl) UpdateTaskByID(ctx context.Context, task *model.Task) (int64, error) {
	t := query.Task
	info, err := t.WithContext(ctx).Where(t.ID.Eq(task.ID)).Updates(map[string]interface{}{
		"status": task.Status,
		"result": task.Result,
	})
	return info.RowsAffected, err
}

func (c *MysqlClientImpl) UpdateGraphByID(ctx context.Context, graph *model.Graph) (int64, error) {
	g := query.Graph
	ret, err := g.WithContext(ctx).Where(g.ID.Eq(graph.ID)).Updates(map[string]interface{}{
		"status":  graph.Status,
		"numNode": graph.NumNode,
		"numEdge": graph.NumEdge,
	})
	return ret.RowsAffected, err
}

func (c *MysqlClientImpl) GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error) {
	g := query.Graph
	graph, err := g.WithContext(ctx).Where(g.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return c.GetGroupByID(ctx, graph.GroupID)
}

func (c *MysqlClientImpl) GetGroupByID(ctx context.Context, id int64) (*model.Group, error) {
	gr := query.Group
	return gr.WithContext(ctx).Where(gr.ID.Eq(id)).Preload(gr.Nodes.NodeAttrs).Preload(gr.Edges.EdgeAttrs).First()
}

func (c *MysqlClientImpl) GetAlgoExecCfgByID(ctx context.Context, id int64) (*model.Algo, error) {
	a := query.Algo
	return a.WithContext(ctx).Select(a.MainClass, a.JarPath).Where(a.ID.Eq(id)).First()
}
