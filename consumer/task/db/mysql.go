package db

import (
	"chainsawman/consumer/task/model"
	"context"
)

type MysqlClient interface {
	UpdateTaskByID(task *model.Task) (int64, error)
	UpdateGraphByID(ctx context.Context, graph *model.Graph) (int64, error)
}
