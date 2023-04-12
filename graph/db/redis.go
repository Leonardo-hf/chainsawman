package db

import (
	"chainsawman/graph/model"
	"context"
)

// RedisClient TODO: 删掉用不上的接口
type RedisClient interface {
	GetTaskById(ctx context.Context, id int64) (*model.KVTask, error)

	UpsertTask(ctx context.Context, task *model.KVTask) error

	ProduceTaskMsg(ctx context.Context, task *model.KVTask) error

	DelTaskMsg(ctx context.Context, id int64) error
}
