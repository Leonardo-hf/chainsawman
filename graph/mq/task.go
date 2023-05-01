package mq

import (
	"chainsawman/graph/model"

	"context"
)

type TaskMq interface {
	ProduceTaskMsg(ctx context.Context, task *model.KVTask) error

	DelTaskMsg(ctx context.Context, id int64) error
}

