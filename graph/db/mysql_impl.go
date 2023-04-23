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

func (c *MysqlClientImpl) InsertTask(task *model.Task) error {
	t := query.Task
	task.Status = 0
	return t.Select(t.Name, t.Params, t.Status).Create(task)
}

func (c *MysqlClientImpl) SearchTaskByID(id int64) (*model.Task, error) {
	t := query.Task
	return t.Where(t.ID.Eq(id)).First()
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
	return g.WithContext(ctx).Where(g.ID.Eq(id)).First()
}

func (c *MysqlClientImpl) GetAllGraph(ctx context.Context) ([]*model.Graph, error) {
	g := query.Graph
	return g.WithContext(ctx).Find()
}
