package db

import (
	"chainsawman/consumer/db/query"
	"chainsawman/consumer/model"
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

func (c *MysqlClientImpl) UpdateTask(task *model.Task) (int64, error) {
	t := query.Task
	info, err := t.Where(t.ID.Eq(task.ID)).Updates(map[string]interface{}{"status": task.Status, "result": task.Result})
	return info.RowsAffected, err
}

func (c *MysqlClientImpl) UpdateGraphStatus(id int64, status int64, ctx context.Context) (int64, error) {
	//TODO implement me
	g := query.Graph
	ret, err := g.WithContext(ctx).Where(g.ID.Eq(id)).Updates(map[string]interface{}{"status": status})
	return ret.RowsAffected, err
}
