package db

import (
	"chainsawman/consumer/task/model"
	"context"
)

type MysqlClient interface {
	GetGroupByGraphId(ctx context.Context, id int64) (*model.Group, error)

	GetGroupByID(ctx context.Context, id int64) (*model.Group, error)

	UpdateAppIDByTaskID(ctx context.Context, task *model.Exec) (int64, error)

	UpdateGraphByID(ctx context.Context, graph *model.Graph) (int64, error)

	GetAlgoExecCfgByID(ctx context.Context, id int64) (*model.Algo, error)
}
