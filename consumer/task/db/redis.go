package db

import (
	"chainsawman/consumer/task/model"

	"context"
)

type RedisClient interface {
	UpsertTask(ctx context.Context, task *model.KVTask) error
}
