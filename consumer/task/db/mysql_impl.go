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

func (c *MysqlClientImpl) UpdateAppIDByTaskID(ctx context.Context, exec *model.Exec) (int64, error) {
	t := query.Exec
	info, err := t.WithContext(ctx).Where(t.ID.Eq(exec.ID)).Update(t.AppID, exec.AppID)
	return info.RowsAffected, err
}

func (c *MysqlClientImpl) UpdateGraphByID(ctx context.Context, graph *model.Graph) (int64, error) {
	g := query.Graph
	ret, err := g.WithContext(ctx).Where(g.ID.Eq(graph.ID)).Updates(map[string]interface{}{
		"status":  graph.Status,
		"numNode": gorm.Expr("numNode + ?", graph.NumNode),
		"numEdge": gorm.Expr("numEdge + ?", graph.NumEdge),
	})
	return ret.RowsAffected, err
}

func (c *MysqlClientImpl) GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error) {
	g := query.Graph
	gr := query.Group
	sub := g.WithContext(ctx).Select(g.GroupID).Where(g.ID.Eq(id))
	group, err := gr.WithContext(ctx).Preload(gr.Nodes.Attrs).Preload(gr.Edges.Attrs).Where(gr.Columns(gr.ID).Eq(sub)).First()
	return group, err
}

func (c *MysqlClientImpl) GetGroupByID(ctx context.Context, id int64) (*model.Group, error) {
	gr := query.Group
	return gr.WithContext(ctx).Where(gr.ID.Eq(id)).Preload(gr.Nodes.Attrs).Preload(gr.Edges.Attrs).First()
}

func (c *MysqlClientImpl) GetAlgoExecCfgByID(ctx context.Context, id int64) (*model.Algo, error) {
	a := query.Algo
	return a.WithContext(ctx).Select(a.MainClass, a.JarPath).Where(a.ID.Eq(id)).First()
}
