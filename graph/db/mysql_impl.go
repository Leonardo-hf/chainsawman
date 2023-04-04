package db

import (
	"chainsawman/graph/db/query"
	"chainsawman/graph/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ClientImpl struct {
}

func InitClientImpl(dsn string) MysqlClient {
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	query.SetDefault(database)
	return &ClientImpl{}
}

func (c *ClientImpl) InsertTask(task *model.Task) error {
	t := query.Task
	task.Status = 0
	return t.Select(t.Name, t.Script, t.Status).Create(task)
}

func (c *ClientImpl) UpdateTask(task *model.Task) (int64, error) {
	t := query.Task
	info, err := t.Update(t.Status, task.Status)
	return info.RowsAffected, err
}

func (c *ClientImpl) SearchTaskById(id int64) (*model.Task, error) {
	t := query.Task
	return t.Where(t.ID.Eq(id)).First()
}
