package db

import (
	"chainsawman/consumer/task/model"
	"time"

	"context"
)

type RedisClient interface {
	UpsertTask(ctx context.Context, task *model.KVTask) error

	DropDuplicate(ctx context.Context, set string, keys []string, expired time.Duration) ([]string, error)
}
