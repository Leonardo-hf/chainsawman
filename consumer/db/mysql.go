package db

import (
	"chainsawman/consumer/model"
)

type MysqlClient interface {
	UpdateTask(task *model.Task) (int64, error)
}
