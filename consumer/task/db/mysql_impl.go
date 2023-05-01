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

func (c *MysqlClientImpl) UpdateTaskByID(task *model.Task) (int64, error) {
	t := query.Task
	info, err := t.Where(t.ID.Eq(task.ID)).Updates(map[string]interface{}{
		"status": task.Status,
		"result": task.Result,
	})
	return info.RowsAffected, err
}

func (c *MysqlClientImpl) UpdateGraphByID(ctx context.Context, graph *model.Graph) (int64, error) {
	g := query.Graph
	ret, err := g.WithContext(ctx).Where(g.ID.Eq(graph.ID)).Updates(map[string]interface{}{
		"status": 1,
		"nodes":  graph.Nodes,
		"edges":  graph.Edges,
	})
	return ret.RowsAffected, err
}
