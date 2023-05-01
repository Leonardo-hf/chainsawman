package mq

import (
	"chainsawman/consumer/task/model"
	"context"
)

type TaskMq interface {
	ConsumeTaskMsg(ctx context.Context, consumer string, handle func(ctx context.Context, task *model.KVTask) error) error
}
