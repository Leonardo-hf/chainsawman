package db

import (
	"chainsawman/graph/model"
	"context"
)

// RedisClient TODO: 删掉用不上的接口
type RedisClient interface {
	GetTaskById(ctx context.Context, id int64) (*model.KVTask, error)

	UpsertTask(ctx context.Context, task *model.KVTask) error

	DropTask(ctx context.Context, id int64) (int64, error)
}
