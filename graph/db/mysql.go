package db

import (
	"chainsawman/graph/model"
	"context"
)

// MysqlClient TODO: 删掉用不上的接口
type MysqlClient interface {
	InsertTask(task *model.Task) error

	SearchTaskById(id int64) (*model.Task, error)

	InsertGraph(graph *model.Graph, ctx context.Context) error

	GetGraphById(id int64, ctx context.Context) (*model.Graph, error)

	GetAllGraph(ctx context.Context) ([]*model.Graph, error)

	UpdateGraphStatus(id int64, status int64, ctx context.Context) (int64, error)
}
