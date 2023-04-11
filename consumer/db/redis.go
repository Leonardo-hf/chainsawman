package db

import (
	"chainsawman/consumer/model"
	"context"
)

type RedisClient interface {
	UpsertTask(ctx context.Context, task *model.KVTask) error

	ConsumeTaskMsg(ctx context.Context, consumer string, handle func(ctx context.Context, task *model.KVTask) error) error
}
