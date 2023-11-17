package db

import (
	"chainsawman/graph/model"
	"context"
)

// MysqlClient
type MysqlClient interface {
	InsertTask(ctx context.Context, task *model.Task) error

	UpdateTaskTIDByID(ctx context.Context, id int64, tid string) (int64, error)

	GetTaskByID(ctx context.Context, id int64) (*model.Task, error)

	InsertGraph(ctx context.Context, graph *model.Graph) error

	DropGraphByID(ctx context.Context, id int64) (int64, error)

	DropTaskByID(ctx context.Context, id int64) (int64, error)

	GetGraphByID(ctx context.Context, id int64) (*model.Graph, error)

	GetGraphByName(ctx context.Context, name string) (*model.Graph, error)

	UpdateGraphStatusByID(ctx context.Context, id int64, status int64) (int64, error)

	GetAllGraph(ctx context.Context) ([]*model.Graph, error)

	GetGraphByGroupID(ctx context.Context, groupID int64) ([]*model.Graph, error)

	GetTasksByGraph(ctx context.Context, graphID int64) ([]*model.Task, error)

	GetTasks(ctx context.Context) ([]*model.Task, error)

	GetAllGroups(ctx context.Context) ([]*model.Group, error)

	GetNodeByID(ctx context.Context, id int64) (*model.Node, error)

	InsertGroup(ctx context.Context, group *model.Group) error

	GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error)

	DropGroupByID(ctx context.Context, id int64) (int64, error)

	GetAllAlgo(ctx context.Context) ([]*model.Algo, error)

	InsertAlgo(ctx context.Context, algo *model.Algo) error

	DropAlgoByID(ctx context.Context, id int64) (int64, error)
}
