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
	return t.WithContext(ctx).Select(t.Idf, t.Params, t.Visible, t.GraphID).Create(task)
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
	return t.WithContext(ctx).Where(t.GraphID.Eq(graphID), t.Visible.Eq(1)).Find()
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
	return g.WithContext(ctx).Where(g.ID.Eq(id)).First()
}

func (c *MysqlClientImpl) GetGraphByName(ctx context.Context, name string) (*model.Graph, error) {
	g := query.Graph
	return g.WithContext(ctx).Where(g.Name.Eq(name)).First()
}

func (c *MysqlClientImpl) GetAllGraph(ctx context.Context) ([]*model.Graph, error) {
	g := query.Graph
	return g.WithContext(ctx).Find()
}
