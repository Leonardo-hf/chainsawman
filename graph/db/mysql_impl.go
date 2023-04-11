package db

import (
	"chainsawman/graph/db/query"
	"chainsawman/graph/model"

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

func (c *MysqlClientImpl) UpdateTask(task *model.Task) (int64, error) {
	t := query.Task
	info, err := t.Update(t.Status, task.Status)
	return info.RowsAffected, err
}

func (c *MysqlClientImpl) SearchTaskById(id int64) (*model.Task, error) {
	t := query.Task
	return t.Where(t.ID.Eq(id)).First()
}
