package db

import (
	"chainsawman/consumer/task/model"
	"context"
)

type MysqlClient interface {
	GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error)

	GetGroupByID(ctx context.Context, id int64) (*model.Group, error)

	UpdateTaskByID(ctx context.Context, task *model.Task) (int64, error)

	UpdateGraphByID(ctx context.Context, graph *model.Graph) (int64, error)
}
