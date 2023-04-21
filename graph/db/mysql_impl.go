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

func (c *MysqlClientImpl) SearchTaskById(id int64) (*model.Task, error) {
	t := query.Task
	return t.Where(t.ID.Eq(id)).First()
}

func (c *MysqlClientImpl) InsertGraph(graph *model.Graph, ctx context.Context) error {
	//TODO implement me
	g := query.Graph
	return g.WithContext(ctx).Create(graph)
}

func (c *MysqlClientImpl) GetGraphById(id int64, ctx context.Context) (*model.Graph, error) {
	//TODO implement me
	g := query.Graph
	return g.WithContext(ctx).Where(g.ID.Eq(id)).First()
}

func (c *MysqlClientImpl) GetAllGraph(ctx context.Context) ([]*model.Graph, error) {
	//TODO implement me
	g := query.Graph
	return g.WithContext(ctx).Find()
}

func (c *MysqlClientImpl) UpdateGraphStatus(id int64, status int64, ctx context.Context) (int64, error) {
	//TODO implement me
	g := query.Graph
	ret, err := g.WithContext(ctx).Where(g.ID.Eq(id)).Updates(map[string]interface{}{"status": status})
	return ret.RowsAffected, err
}
