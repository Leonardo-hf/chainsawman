package db

import (
	"chainsawman/graph/model"
	"context"
)

// MysqlClient TODO: 删掉用不上的接口
type MysqlClient interface {
	InsertTask(task *model.Task) error

	SearchTaskByID(id int64) (*model.Task, error)

	InsertGraph(ctx context.Context, graph *model.Graph) error

	DropGraphByID(ctx context.Context, id int64) (int64, error)

	GetGraphByID(ctx context.Context, id int64) (*model.Graph, error)

	GetAllGraph(ctx context.Context) ([]*model.Graph, error)
}
