package db

import (
	"chainsawman/graph/model"
	"context"
)

// MysqlClient
type MysqlClient interface {
	InsertTask(ctx context.Context, task *model.Task) error

	GetTaskByID(ctx context.Context, id int64) (*model.Task, error)

	InsertGraph(ctx context.Context, graph *model.Graph) error

	DropGraphByID(ctx context.Context, id int64) (int64, error)

	GetGraphByID(ctx context.Context, id int64) (*model.Graph, error)

	GetGraphByName(ctx context.Context, name string) (*model.Graph, error)

	GetAllGraph(ctx context.Context) ([]*model.Graph, error)

	GetTasksByGraph(ctx context.Context, graphID int64) ([]*model.Task, error)
}
