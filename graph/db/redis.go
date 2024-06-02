package db

import (
	"chainsawman/graph/model"
	"context"
	"time"
)

type RedisClient interface {
	GetTaskById(ctx context.Context, id string) (*model.KVTask, error)

	UpsertTask(ctx context.Context, task *model.KVTask) error

	DropTask(ctx context.Context, id string) (int64, error)

	CheckIdempotent(ctx context.Context, id string, expire time.Duration) (bool, error)
}
