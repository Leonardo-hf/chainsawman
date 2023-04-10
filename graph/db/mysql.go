package db

import (
	"chainsawman/graph/model"
)

// MysqlClient TODO: 删掉用不上的接口
type MysqlClient interface {
	InsertTask(task *model.Task) error

	UpdateTask(task *model.Task) (int64, error)

	SearchTaskById(id int64) (*model.Task, error)
}
