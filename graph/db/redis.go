package db

import (
	"chainsawman/graph/model"
	"context"
)

type RedisClient interface {
	GetTaskById(ctx context.Context, id int64) (*model.KVTask, error)

	UpsertTask(ctx context.Context, task *model.KVTask) error

	ProduceTaskMsg(ctx context.Context, task *model.KVTask) error

	DelTaskMsg(ctx context.Context, id int64) error

	ConsumeTaskMsg(ctx context.Context, consumer string, handle func(task *model.KVTask) error) error
}
