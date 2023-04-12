package db

import (
	"chainsawman/graph/model"
)

// MysqlClient TODO: 删掉用不上的接口
type MysqlClient interface {
	InsertTask(task *model.Task) error

	SearchTaskById(id int64) (*model.Task, error)
}
