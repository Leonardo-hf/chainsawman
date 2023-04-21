package db

import (
	"chainsawman/consumer/model"
	"context"
)

type MysqlClient interface {
	UpdateTask(task *model.Task) (int64, error)
	UpdateGraphStatus(id int64, status int64, ctx context.Context) (int64, error)
}
