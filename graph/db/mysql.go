package db

import (
	"chainsawman/graph/model"
	"context"
	"time"
)

// MysqlClient
type MysqlClient interface {
	GetLatestSETask(ctx context.Context, n int64) ([]*SETask, error)

	GetLatestHHITask(ctx context.Context, n int64) ([]*HHITask, error)

	InsertAlgoTask(ctx context.Context, exec *model.Exec) error

	GetAlgoTasks(ctx context.Context) ([]*model.Exec, error)

	GetAlgoTasksByGraphId(ctx context.Context, id int64) ([]*model.Exec, error)

	GetAlgoTaskByID(ctx context.Context, id int64) (*model.Exec, error)

	DropAlgoTaskByID(ctx context.Context, id int64) (int64, error)

	UpdateAlgoTaskTIDByID(ctx context.Context, id int64, tid string) (int64, error)

	InsertGraph(ctx context.Context, graph *model.Graph) error

	DropGraphByID(ctx context.Context, id int64) (int64, error)

	GetGraphByID(ctx context.Context, id int64) (*model.Graph, error)

	GetGraphByName(ctx context.Context, name string) (*model.Graph, error)

	UpdateGraphStatusByID(ctx context.Context, id int64, status int64) (int64, error)

	GetAllGraph(ctx context.Context) ([]*model.Graph, error)

	GetGraphByGroupID(ctx context.Context, groupID int64) ([]*model.Graph, error)

	GetAllGroups(ctx context.Context) ([]*model.Group, error)

	GetNodeByID(ctx context.Context, id int64) (*model.Node, error)

	InsertGroup(ctx context.Context, group *model.Group) error

	GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error)

	DropGroupByID(ctx context.Context, id int64) (int64, error)

	GetAllAlgo(ctx context.Context) ([]*model.Algo, error)

	GetAlgoExecCfgByID(ctx context.Context, id int64) (*model.Algo, error)

	InsertAlgo(ctx context.Context, algo *model.Algo) error

	DropAlgoByID(ctx context.Context, id int64) (int64, error)
}

type HHITask struct {
	Output     string
	UpdateTime time.Time `gorm:"column:updateTime;"`
	Graph      string
}

type SETask struct {
	Output     string
	UpdateTime time.Time `gorm:"column:updateTime;"`
	Graph      string
	GraphID    int64
	Algo       string
}
